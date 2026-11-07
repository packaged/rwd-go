package country

import "testing"

func TestFromLocale(t *testing.T) {
	for in, want := range map[string]string{
		"en-GB": "GB", "en_gb": "GB", "fil-PH": "PH", "zh-Hant-TW": "TW", "sr-Latn-RS": "RS", "pt-BR": "BR",
	} {
		if c, ok := FromLocale(in); !ok || c.GetIso2() != want {
			t.Errorf("FromLocale(%q) = %v, %v; want %s", in, c, ok, want)
		}
	}
	for _, in := range []string{"en", "zh-Hant", "es-419", "", "not a tag!"} {
		if c, ok := FromLocale(in); ok {
			t.Errorf("FromLocale(%q) = %s, want none", in, c.GetIso2())
		}
	}
}
