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

func TestRegion(t *testing.T) {
	for in, want := range map[string]string{
		"en-GB": "GB", "en_gb": "GB", "fil-PH": "PH", "zh-Hant-TW": "TW", "sr-Latn-RS": "RS",
		"es-419": "419", "en": "", "zh-Hant": "", "": "", "not a tag!": "",
	} {
		if got := Region(in); got != want {
			t.Errorf("Region(%q) = %q, want %q", in, got, want)
		}
	}
}

// RFC 5646 section 4.4.2, figure 8: each step drops the right-most subtag, and a
// single-character subtag left at the end goes with it.
func TestTruncateFollowsRFC5646(t *testing.T) {
	const tag = "zh-Latn-CN-variant1-a-extend1-x-wadegile-private1"
	steps := []struct {
		maxLen int
		want   string
	}{
		{len(tag), tag},
		{len(tag) - 1, "zh-Latn-CN-variant1-a-extend1-x-wadegile"},
		{len("zh-Latn-CN-variant1-a-extend1-x-wadegile") - 1, "zh-Latn-CN-variant1-a-extend1"},
		{len("zh-Latn-CN-variant1-a-extend1") - 1, "zh-Latn-CN-variant1"},
		{len("zh-Latn-CN-variant1") - 1, "zh-Latn-CN"},
		{len("zh-Latn-CN") - 1, "zh-Latn"},
		{len("zh-Latn") - 1, "zh"},
	}
	for _, s := range steps {
		if got := Truncate(tag, s.maxLen); got != s.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tag, s.maxLen, got, s.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		tag    string
		maxLen int
		want   string
	}{
		{"en-GB", 8, "en-GB"},
		{"en-GB-oxendict", 8, "en-GB"},
		{"zh-Hans-CN", 8, "zh-Hans"},
		{"sr-Latn-RS", 8, "sr-Latn"},
		{"en-US-u-ca-gregory", 8, "en-US"},
		{"  en-GB-oxendict  ", 8, "en-GB"},
		{"abcdefghij", 8, ""},
		{"", 8, ""},
	}
	for _, tt := range tests {
		if got := Truncate(tt.tag, tt.maxLen); got != tt.want || len(got) > tt.maxLen {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.tag, tt.maxLen, got, tt.want)
		}
	}
}
