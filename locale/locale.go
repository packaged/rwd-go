// Package locale reads BCP 47 language tags ("en-GB", "zh-Hant-TW", "fil-PH") for the
// language and region they name.
package locale

import (
	"regexp"
	"strings"

	helpers "github.com/packaged/helpers-go"
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

// maxBrowserLanguageLength is the longest browser language tag we emit.
const maxBrowserLanguageLength = 8

// browserLanguageRegion matches a BCP 47 region subtag: two letters, or the
// three-digit UN M.49 form ("419" for Latin America). Script subtags are always
// four letters and variants five to eight, so neither can be mistaken for one.
var browserLanguageRegion = regexp.MustCompile(`^([A-Za-z]{2}|[0-9]{3})$`)

// NormalizeBrowserLanguage shortens a browser-supplied BCP 47 language tag to
// fit maxBrowserLanguageLength. Tags carrying a script or variant subtag
// ("en-GB-oxendict", "zh-Hans-CN", "ca-ES-valencia") are perfectly valid but
// longer than that.
//
// It reduces to language-region, dropping any script and variant in between,
// rather than cutting subtags off the end. The two differ only for tags carrying
// both a script and a region, where cutting from the end keeps the script
// ("zh-Hans") and this keeps the region ("zh-CN"). language-region is the shape
// browsers overwhelmingly report and consumers therefore recognise, and for
// Chinese the region implies the script anyway.
//
// The result is always a well-formed tag - truncating to length would emit
// "en-GB-ox".
func NormalizeBrowserLanguage(tag string) string {
	tag = strings.TrimSpace(tag)
	if len(tag) <= maxBrowserLanguageLength {
		return tag
	}

	subtags := strings.Split(tag, "-")
	language := subtags[0]

	for _, subtag := range subtags[1:] {
		// A single-character subtag is an extension singleton; no region follows.
		if len(subtag) == 1 {
			break
		}
		if browserLanguageRegion.MatchString(subtag) {
			if candidate := language + "-" + subtag; len(candidate) <= maxBrowserLanguageLength {
				return candidate
			}
			break
		}
	}

	// An oversized primary subtag is not a valid language tag whatever we do, so
	// keep the result within length and let the consumer judge the value.
	return helpers.TruncateStringToRune(language, maxBrowserLanguageLength)
}
