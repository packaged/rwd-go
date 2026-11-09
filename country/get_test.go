package country

import (
	"testing"

	"golang.org/x/text/language"
)

func TestKosovo(t *testing.T) {
	c, ok := List["XK"]
	if !ok {
		t.Fatal("List has no XK")
	}
	if c.GetName() != "Kosovo" || c.GetIso2() != "XK" || c.GetIso3() != "XKX" ||
		c.GetIsoNumeric() != "" || c.GetDialPrefix() != 383 || c.GetCurrencyCode() != "EUR" {
		t.Errorf("XK = %s %s %s %q %d %s, want Kosovo XK XKX \"\" 383 EUR",
			c.GetName(), c.GetIso2(), c.GetIso3(), c.GetIsoNumeric(), c.GetDialPrefix(), c.GetCurrencyCode())
	}
	for locale, want := range map[string]string{"en": "Kosovo", "de": "Kosovo", "ja": "コソボ", "ru": "Косово", "pl": "Kosowo"} {
		if got := c.GetTranslatedName(locale); got != want {
			t.Errorf("GetTranslatedName(%q) = %q, want %q", locale, got, want)
		}
	}
	if codes := DialCodes[383]; len(codes) != 1 || codes[0] != "XK" {
		t.Errorf("DialCodes[383] = %v, want [XK]", codes)
	}
	for _, locale := range SupportedLocales {
		if got := getTranslator(language.Make(locale)).Translate("country.XK"); got == "" {
			t.Errorf("%s has no translation for country.XK", locale)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		code, want string
		ok         bool
	}{
		{"GB", "GB", true},
		{"gb", "GB", true},
		{"UK", "GB", true}, // exceptionally reserved for the United Kingdom
		{"FX", "FR", true}, // exceptionally reserved for metropolitan France
		{"XK", "XK", true},
		{"ZZ", "", false},
		{"", "", false},
	}
	for _, tc := range tests {
		c, ok := Get(tc.code)
		if ok != tc.ok || (ok && c.GetIso2() != tc.want) {
			t.Errorf("Get(%q) = %v %v, want %q %v", tc.code, c, ok, tc.want, tc.ok)
		}
	}
}

// Aliases resolve through Get only; List stays one entry per country, as callers
// render it as a country picker.
func TestAliasesAreNotInList(t *testing.T) {
	for _, code := range []string{"UK", "FX"} {
		if _, ok := List[code]; ok {
			t.Errorf("List has alias %s", code)
		}
	}
}
