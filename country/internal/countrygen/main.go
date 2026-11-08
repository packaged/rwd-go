// Command countrygen writes the country data from its sources: ISO 3166-1 codes, English
// names, translations, each country's currency and EU membership from Unicode CLDR, and
// dialling codes from Google's libphonenumber.
//
//	go run ./internal/countrygen -cldr release-49 -libphonenumber v9.0.43
//
// It writes countries.go and translations/*.yaml in the current directory, which is the
// country package.
package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// locales are the translations written, matching SupportedLocales.
var locales = []string{
	"ar", "bg", "cs", "da", "de", "el", "en", "es", "et", "fi",
	"fr", "hr", "hu", "it", "ja", "ko", "lt", "lv", "nl", "pl",
	"pt", "ro", "ru", "sk", "sl", "sv", "tr", "zh",
}

// kosovo is the one country without an ISO 3166-1 code. XK is the user-assigned code
// the European Commission, the IMF and the card schemes use, with XKX as its alpha-3;
// it has no numeric code.
const kosovo = "XK"

func main() {
	cldr := flag.String("cldr", "", "CLDR release tag, such as release-49")
	phone := flag.String("libphonenumber", "", "libphonenumber release tag, such as v9.0.43")
	flag.Parse()
	if *cldr == "" || *phone == "" {
		fail(fmt.Errorf("-cldr and -libphonenumber are required"))
	}
	cldrBase := "https://raw.githubusercontent.com/unicode-org/cldr/" + *cldr + "/common/"
	phoneURL := "https://raw.githubusercontent.com/google/libphonenumber/" + *phone + "/resources/PhoneNumberMetadata.xml"

	var validity struct {
		IDs []struct {
			Status string `xml:"idStatus,attr"`
			Type   string `xml:"type,attr"`
			Codes  string `xml:",chardata"`
		} `xml:"idValidity>id"`
	}
	fetchXML(cldrBase+"validity/region.xml", &validity)
	regular := map[string]bool{}
	for _, id := range validity.IDs {
		if id.Type == "region" && id.Status == "regular" {
			for _, code := range expand(id.Codes) {
				regular[code] = true
			}
		}
	}

	var supplemental struct {
		Codes []struct {
			Type    string `xml:"type,attr"`
			Numeric string `xml:"numeric,attr"`
			Alpha3  string `xml:"alpha3,attr"`
		} `xml:"codeMappings>territoryCodes"`
		Regions []struct {
			Code       string `xml:"iso3166,attr"`
			Currencies []struct {
				Code   string `xml:"iso4217,attr"`
				To     string `xml:"to,attr"`
				Tender string `xml:"tender,attr"`
			} `xml:"currency"`
		} `xml:"currencyData>region"`
		Groups []struct {
			Type     string `xml:"type,attr"`
			Contains string `xml:"contains,attr"`
			Status   string `xml:"status,attr"`
		} `xml:"territoryContainment>group"`
	}
	fetchXML(cldrBase+"supplemental/supplementalData.xml", &supplemental)

	// An ISO 3166-1 country is a regular CLDR region with an ISO numeric code. The
	// regions CLDR adds beyond ISO 3166-1, such as Ascension Island (AC) and the Canary
	// Islands (IC), have none.
	countries := map[string]*country{}
	for _, c := range supplemental.Codes {
		if !regular[c.Type] || len(c.Type) != 2 || (c.Numeric == "" && c.Type != kosovo) {
			continue
		}
		cn := &country{Iso2: c.Type, Iso3: c.Alpha3}
		if c.Type == kosovo {
			cn.Iso3 = "XKX"
		} else {
			n, err := strconv.ParseUint(c.Numeric, 10, 16)
			if err != nil {
				fail(fmt.Errorf("numeric code for %s: %w", c.Type, err))
			}
			cn.Numeric = uint16(n)
		}
		countries[c.Type] = cn
	}

	// CLDR lists a region's currencies most recent first; the first still in use as
	// legal tender is the country's currency.
	for _, r := range supplemental.Regions {
		cn, ok := countries[r.Code]
		if !ok {
			continue
		}
		for _, cur := range r.Currencies {
			if cur.To == "" && cur.Tender != "false" {
				cn.Currency = cur.Code
				break
			}
		}
	}

	var eu []string
	for _, g := range supplemental.Groups {
		if g.Type == "EU" && g.Status == "" {
			eu = strings.Fields(g.Contains)
		}
	}
	if len(eu) == 0 {
		fail(fmt.Errorf("CLDR has no EU group"))
	}
	sort.Strings(eu)

	var metadata struct {
		Territories []struct {
			ID          string `xml:"id,attr"`
			CountryCode uint16 `xml:"countryCode,attr"`
		} `xml:"territories>territory"`
	}
	fetchXML(phoneURL, &metadata)
	for _, t := range metadata.Territories {
		if cn, ok := countries[t.ID]; ok {
			cn.Dial = t.CountryCode
		}
	}

	names := map[string]map[string]string{}
	var alternates map[string][]string
	for _, loc := range locales {
		var ldml struct {
			Territories []struct {
				Type string `xml:"type,attr"`
				Alt  string `xml:"alt,attr"`
				Name string `xml:",chardata"`
			} `xml:"localeDisplayNames>territories>territory"`
		}
		fetchXML(cldrBase+"main/"+loc+".xml", &ldml)
		names[loc] = map[string]string{}
		if loc == "en" {
			alternates = map[string][]string{}
		}
		for _, t := range ldml.Territories {
			if _, ok := countries[t.Type]; !ok {
				continue
			}
			switch {
			case t.Alt == "":
				names[loc][t.Type] = t.Name
			case loc == "en" && (t.Alt == "short" || t.Alt == "variant"):
				alternates[t.Type] = append(alternates[t.Type], t.Name)
			}
		}
		for code := range countries {
			if names[loc][code] == "" {
				fail(fmt.Errorf("CLDR %s has no name for %s", loc, code))
			}
		}
	}
	for code, cn := range countries {
		cn.Name = names["en"][code]
	}

	writeGo(*cldr, *phone, countries, alternates, eu)
	for _, loc := range locales {
		writeTranslations(*cldr, loc, names[loc])
	}
}

type country struct {
	Name, Iso2, Iso3, Currency string
	Numeric, Dial              uint16
}

func writeGo(cldr, phone string, countries map[string]*country, alternates map[string][]string, eu []string) {
	codes := make([]string, 0, len(countries))
	for code := range countries {
		codes = append(codes, code)
	}
	sort.Strings(codes)

	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by internal/countrygen from Unicode CLDR %s and libphonenumber %s; DO NOT EDIT.\n\n", cldr, phone)
	b.WriteString("package country\n\nvar List map[string]Country\n\n")
	b.WriteString("// DialCodes lists the countries using each international dialling code.\nvar DialCodes map[uint16][]string\n\n")
	b.WriteString("// EUCountryCodes are the member states of the European Union.\n")
	fmt.Fprintf(&b, "var EUCountryCodes = %#v\n\n", eu)
	b.WriteString("// alternateNames are the other English names CLDR gives a country.\nvar alternateNames = map[string][]string{\n")
	for _, code := range codes {
		if alt := alternates[code]; len(alt) > 0 {
			fmt.Fprintf(&b, "%q: %#v,\n", code, alt)
		}
	}
	b.WriteString("}\n\nfunc init() {\nList = map[string]Country{\n")
	dial := map[uint16][]string{}
	for _, code := range codes {
		c := countries[code]
		fmt.Fprintf(&b, "%q: CommonCountry{name: %q, iso2: %q, iso3: %q, numericCode: %d, dialPrefix: %d, currencyCode: %q},\n",
			code, c.Name, c.Iso2, c.Iso3, c.Numeric, c.Dial, c.Currency)
		if c.Dial != 0 {
			dial[c.Dial] = append(dial[c.Dial], code)
		}
	}
	b.WriteString("}\nDialCodes = map[uint16][]string{\n")
	prefixes := make([]uint16, 0, len(dial))
	for p := range dial {
		prefixes = append(prefixes, p)
	}
	slices.Sort(prefixes)
	for _, p := range prefixes {
		fmt.Fprintf(&b, "%d: %#v,\n", p, dial[p])
	}
	b.WriteString("}\n}\n")

	src, err := format.Source(b.Bytes())
	if err != nil {
		fail(err)
	}
	if err := os.WriteFile("countries.go", src, 0o644); err != nil {
		fail(err)
	}
}

func writeTranslations(cldr, loc string, names map[string]string) {
	codes := make([]string, 0, len(names))
	for code := range names {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Country names from Unicode CLDR %s. Generated by internal/countrygen; do not edit.\n", cldr)
	for _, code := range codes {
		// A JSON string is a valid YAML double-quoted scalar.
		v, _ := json.Marshal(names[code])
		fmt.Fprintf(&b, "country.%s: %s\n", code, v)
	}
	if err := os.WriteFile(filepath.Join("translations", loc+".yaml"), b.Bytes(), 0o644); err != nil {
		fail(err)
	}
}

// expand reads CLDR's validity list: codes separated by space, where "AC~G" stands for
// AC to AG, the last character running through the range.
func expand(list string) []string {
	var out []string
	for _, f := range strings.Fields(list) {
		start, end, ranged := strings.Cut(f, "~")
		if !ranged {
			out = append(out, f)
			continue
		}
		prefix, from := start[:len(start)-len(end)], start[len(start)-len(end):]
		for c := from[0]; c <= end[0]; c++ {
			out = append(out, prefix+string(c)+from[1:])
		}
	}
	return out
}

func fetchXML(url string, v any) {
	resp, err := http.Get(url)
	if err != nil {
		fail(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fail(fmt.Errorf("%s: %s", url, resp.Status))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fail(err)
	}
	if err := xml.Unmarshal(body, v); err != nil {
		fail(fmt.Errorf("%s: %w", url, err))
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "countrygen:", err)
	os.Exit(1)
}
