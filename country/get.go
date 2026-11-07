package country

import (
	"strconv"
	"strings"
)

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

// GetByIso3 returns the country for an ISO 3166-1 alpha-3 code, case-insensitive.
func GetByIso3(code string) (Country, bool) {
	code = strings.ToUpper(strings.TrimSpace(code))
	for _, c := range List {
		if c.GetIso3() == code {
			return c, true
		}
	}
	return nil, false
}

// GetByNumericCode returns the country for an ISO 3166-1 numeric code. 0 finds no
// country, including those without a numeric code, such as Kosovo.
func GetByNumericCode(code uint16) (Country, bool) {
	if code == 0 {
		return nil, false
	}
	for _, c := range List {
		if c.GetNumericCode() == code {
			return c, true
		}
	}
	return nil, false
}

// GetByName returns the country whose English name matches, ignoring case and
// surrounding space.
func GetByName(name string) (Country, bool) {
	name = strings.TrimSpace(name)
	for _, c := range List {
		if strings.EqualFold(c.GetName(), name) {
			return c, true
		}
	}
	return nil, false
}

// Lookup returns the country for an ISO 3166-1 code in any of its forms, case-insensitive:
// alpha-2 (including UK and FX), alpha-3, or numeric digits such as "826".
func Lookup(code string) (Country, bool) {
	code = strings.TrimSpace(code)
	if code != "" && strings.Trim(code, "0123456789") == "" {
		n, err := strconv.ParseUint(code, 10, 16)
		if err != nil {
			return nil, false
		}
		return GetByNumericCode(uint16(n))
	}
	switch len(code) {
	case 2:
		return Get(code)
	case 3:
		return GetByIso3(code)
	}
	return nil, false
}
