package country

import "github.com/packaged/rwd-go/locale"

// FromLocale returns the country a BCP 47 tag names by its region: Taiwan for
// "zh-Hant-TW". A tag without a region ("en") or with one that is not a country
// ("es-419") has none; the country is never inferred from the language.
func FromLocale(tag string) (Country, bool) {
	region := locale.Region(tag)
	if len(region) != 2 {
		return nil, false
	}
	return Get(region)
}
