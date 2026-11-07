// Package locale reads BCP 47 language tags ("en-GB", "zh-Hant-TW", "fil-PH") for the
// language and region they name.
package locale

import (
	"strings"

	"golang.org/x/text/language"
)

// Language returns the ISO 639 language of a BCP 47 tag ("fil" for "fil-PH", "en" for
// "en_GB"), or "" when the tag does not parse.
func Language(tag string) string {
	t, err := language.Parse(strings.TrimSpace(tag))
	if err != nil {
		return ""
	}
	base, _ := t.Base()
	return base.String()
}

// Region returns the region subtag a BCP 47 tag names explicitly, upper-cased: "TW" for
// "zh-Hant-TW", "419" for "es-419", and "" for a tag without one, such as "en". It is
// never inferred from the language.
func Region(tag string) string {
	t, err := language.Parse(strings.TrimSpace(tag))
	if err != nil {
		return ""
	}
	_, _, region := t.Raw()
	if region.String() == "ZZ" {
		return ""
	}
	return region.String()
}

// Truncate shortens a BCP 47 tag to at most maxLen characters as RFC 5646 section 4.4.2
// requires: subtags are removed from the right, one at a time, and a single-character
// subtag left at the end is removed with them. "zh-Latn-CN-variant1-a-extend1" truncated
// to 10 is "zh-Latn-CN". A tag already within maxLen is returned unchanged, and one whose
// primary language subtag alone is longer than maxLen is returned as "".
func Truncate(tag string, maxLen int) string {
	tag = strings.TrimSpace(tag)
	for len(tag) > maxLen {
		i := strings.LastIndexByte(tag, '-')
		if i < 0 {
			return ""
		}
		tag = tag[:i]
		if j := strings.LastIndexByte(tag, '-'); j >= 0 && len(tag)-j == 2 {
			tag = tag[:j]
		}
	}
	return tag
}
