package country

import (
	"embed"
	"strconv"
	"strings"
	"sync"

	"github.com/packaged/i18n-go"
	"golang.org/x/text/language"
)

//go:embed translations/*
var translationsFs embed.FS

var SupportedLocales = []string{
	"ar", "bg", "cs", "da", "de", "el", "en", "es", "et", "fi",
	"fr", "hr", "hu", "it", "ja", "ko", "lt", "lv", "nl", "pl",
	"pt", "ro", "ru", "sk", "sl", "sv", "tr", "zh",
}

var (
	catalogOnce sync.Once
	catalog     *i18n.Catalog
)

func buildCatalog() {
	catalog = i18n.NewEmptyCatalog()
	for _, loc := range SupportedLocales {
		catalog.Quick(language.Make(loc)).Add(
			i18n.NewEmbeddedFile(translationsFs, "translations/"+loc+".yaml"),
		)
	}
}

type Country interface {
	GetName() string
	GetIso2() string
	GetIso3() string
	// GetIsoNumeric returns the ISO 3166-1 numeric code as its three digits: "056" for
	// Belgium, or "" for a country without one, such as Kosovo.
	GetIsoNumeric() string
	// Deprecated: use GetIsoNumeric. GetNumericCode is the code's value, without the
	// leading zeros the code is written with (56 for Belgium's "056").
	GetNumericCode() uint16
	GetDialPrefix() uint16
	GetCurrencyCode() string
	GetTranslatedName(locale string) string
}

type CommonCountry struct {
	name         string
	iso2         string
	iso3         string
	numericCode  uint16
	dialPrefix   uint16
	currencyCode string
}

func (c CommonCountry) GetName() string {
	return c.name
}

func (c CommonCountry) GetIso2() string {
	return c.iso2
}

func (c CommonCountry) GetIso3() string {
	return c.iso3
}

func (c CommonCountry) GetIsoNumeric() string {
	return isoNumeric(c.numericCode)
}

// isoNumeric writes an ISO numeric code as its three digits, or "" for 0, which no
// country or currency has.
func isoNumeric(code uint16) string {
	if code == 0 {
		return ""
	}
	s := strconv.Itoa(int(code))
	return strings.Repeat("0", 3-len(s)) + s
}

// Deprecated: use GetIsoNumeric. GetNumericCode is the code's value, without the
// leading zeros the code is written with (56 for Belgium's "056").
func (c CommonCountry) GetNumericCode() uint16 {
	return c.numericCode
}

func (c CommonCountry) GetDialPrefix() uint16 {
	return c.dialPrefix
}

func (c CommonCountry) GetCurrencyCode() string {
	return c.currencyCode
}

func (c CommonCountry) GetTranslatedName(locale string) string {
	translator := getTranslator(language.Make(locale))
	if translator == nil {
		return c.name
	}
	if name := translator.Translate("country." + c.iso2); name != "" {
		return name
	}
	return c.name
}

func getTranslator(locale language.Tag) i18n.Translator {
	catalogOnce.Do(buildCatalog)
	return catalog.GetTranslator(locale)
}
