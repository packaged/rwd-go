package country

import "strings"

// aliases are the exceptionally reserved ISO 3166-1 codes some data sources use for a
// country List holds under its own code.
var aliases = map[string]string{
	"UK": "GB", // United Kingdom
	"FX": "FR", // metropolitan France
}

// Get returns the country for an ISO 3166-1 alpha-2 code in either case, resolving the
// exceptionally reserved codes UK and FX, which List does not hold.
func Get(code string) (Country, bool) {
	code = strings.ToUpper(code)
	if iso2, ok := aliases[code]; ok {
		code = iso2
	}
	c, ok := List[code]
	return c, ok
}
