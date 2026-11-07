// Package subdivision holds ISO 3166-2 country subdivisions (states, provinces and
// territories) and finds them by code or name.
package subdivision

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Subdivision is an ISO 3166-2 country subdivision.
type Subdivision struct {
	// Code is the full ISO 3166-2 code, such as "US-FL".
	Code string
	// Name is the English name, such as "Florida".
	Name string
	// aliases are other names the subdivision is known by, such as "Québec".
	aliases []string
}

// Country returns the ISO 3166-1 alpha-2 code of the subdivision's country: "US" for
// "US-FL".
func (s Subdivision) Country() string {
	country, _, _ := strings.Cut(s.Code, "-")
	return country
}

// LocalCode returns the code within its country: "FL" for "US-FL".
func (s Subdivision) LocalCode() string {
	_, local, _ := strings.Cut(s.Code, "-")
	return local
}

// List returns the subdivisions of a country, by its ISO 3166-1 alpha-2 code,
// case-insensitive. A country without subdivision data has none.
func List(country string) []Subdivision {
	return byCountry[strings.ToUpper(strings.TrimSpace(country))]
}

// Lookup finds a country's subdivision by its local code ("FL"), full code ("US-FL") or
// name ("Florida"), ignoring case, accents, spacing and punctuation, so "Québec",
// "quebec" and "QC" all find Quebec. The country is an ISO 3166-1 alpha-2 code,
// case-insensitive.
func Lookup(country, codeOrName string) (Subdivision, bool) {
	country = strings.ToUpper(strings.TrimSpace(country))
	key := fold(codeOrName)
	if key == "" {
		return Subdivision{}, false
	}
	for _, s := range byCountry[country] {
		if key == fold(s.LocalCode()) || key == fold(s.Code) || key == fold(s.Name) {
			return s, true
		}
		for _, alias := range s.aliases {
			if key == fold(alias) {
				return s, true
			}
		}
	}
	return Subdivision{}, false
}

// fold reduces a code or name to lower-case letters and digits, without accents.
func fold(s string) string {
	var b strings.Builder
	for _, r := range norm.NFD.String(s) {
		switch {
		case unicode.Is(unicode.Mn, r):
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
