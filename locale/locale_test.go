package locale

import "testing"

func TestLanguage(t *testing.T) {
	for in, want := range map[string]string{
		"en": "en", "en-GB": "en", "en_GB": "en", "fil-PH": "fil", "haw-US": "haw",
		"zh-Hant-TW": "zh", " pt-BR ": "pt", "": "", "not a tag!": "",
	} {
		if got := Language(in); got != want {
			t.Errorf("Language(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCountry(t *testing.T) {
	for in, want := range map[string]string{
		"en-GB": "GB", "en_gb": "GB", "fil-PH": "PH", "zh-Hant-TW": "TW", "sr-Latn-RS": "RS", "pt-BR": "BR",
	} {
		if c, ok := Country(in); !ok || c.GetIso2() != want {
			t.Errorf("Country(%q) = %v, %v; want %s", in, c, ok, want)
		}
	}
	for _, in := range []string{"en", "zh-Hant", "es-419", "", "not a tag!"} {
		if c, ok := Country(in); ok {
			t.Errorf("Country(%q) = %s, want none", in, c.GetIso2())
		}
	}
}

func TestNormalizeBrowserLanguage(t *testing.T) {
	tests := []struct{ name, tag, expect string }{
		{"empty", "", ""},
		{"language only", "en", "en"},
		{"language and region", "en-GB", "en-GB"},
		{"region preserved", "pt-BR", "pt-BR"},
		{"script kept when it fits alone", "zh-Hans", "zh-Hans"},
		{"numeric region kept", "es-419", "es-419"},
		{"at the limit", "ca-ES-va", "ca-ES-va"},
		{"variant dropped", "en-GB-oxendict", "en-GB"},
		{"valencian variant dropped", "ca-ES-valencia", "ca-ES"},
		{"script dropped, region kept", "zh-Hans-CN", "zh-CN"},
		{"traditional chinese region kept", "zh-Hant-TW", "zh-TW"},
		{"hong kong region kept", "zh-Hant-HK", "zh-HK"},
		{"serbian script dropped for region", "sr-Latn-RS", "sr-RS"},
		{"azerbaijani script dropped for region", "az-Latn-AZ", "az-AZ"},
		{"year variant dropped", "de-DE-1996", "de-DE"},
		{"extlang skipped for region", "zh-yue-HK", "zh-HK"},
		{"extension singleton stops the scan", "en-US-u-ca-gregory", "en-US"},
		{"no region falls back to language", "abcdefg-Hant", "abcdefg"},
		{"oversized primary subtag truncated", "abcdefghij", "abcdefgh"},
		{"region that would not fit is dropped", "abcdefgh-GB", "abcdefgh"},
		{"whitespace trimmed", "  en-GB-oxendict  ", "en-GB"},
		{"lowercase region matched", "en-gb-oxendict", "en-gb"},
	}
	for _, tt := range tests {
		got := NormalizeBrowserLanguage(tt.tag)
		if got != tt.expect || len(got) > maxBrowserLanguageLength {
			t.Errorf("%s: NormalizeBrowserLanguage(%q) = %q, want %q", tt.name, tt.tag, got, tt.expect)
		}
	}
}
