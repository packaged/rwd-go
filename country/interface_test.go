package country

import (
	"testing"

	"golang.org/x/text/language"
)

func TestGetCatalog(t *testing.T) {
	if translator := getTranslator(language.English); translator == nil {
		t.Fatal("expected English translator, got nil")
	}
}

func TestGetTranslatedName(t *testing.T) {
	tests := []struct {
		name   string
		iso2   string
		locale string
		want   string
	}{
		{"english US", "US", "en", "United States"},
		{"english GB", "GB", "en", "United Kingdom"},
		{"english FR", "FR", "en", "France"},
		{"english AF", "AF", "en", "Afghanistan"},
		{"missing key returns empty", "ZZ", "en", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := CommonCountry{iso2: tc.iso2}
			if got := c.GetTranslatedName(tc.locale); got != tc.want {
				t.Errorf("GetTranslatedName(%q) for %q = %q, want %q", tc.locale, tc.iso2, got, tc.want)
			}
		})
	}
}

func TestGetTranslatedNameFromList(t *testing.T) {
	c, ok := List["US"]
	if !ok {
		t.Fatal("expected US in List")
	}

	if got := c.GetTranslatedName("en"); got != "United States" {
		t.Errorf("List[\"US\"].GetTranslatedName(\"en\") = %q, want %q", got, "United States")
	}
}
