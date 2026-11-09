package currency

import (
	"fmt"
	"strings"
)

// GetByCode returns the currency for an ISO 4217 alphabetic code, case-insensitive, and
// whether it was found.
func GetByCode(code string) (Currency, bool) {
	cur, ok := Currencies[strings.ToUpper(strings.TrimSpace(code))]
	return cur, ok
}

// GetByNumericCode returns the currency for the value of an ISO 4217 numeric code (8 for
// the lek's "008"), and whether it was found.
func GetByNumericCode(code uint16) (Currency, bool) {
	if code == 0 || code > 999 {
		return nil, false
	}
	iso := fmt.Sprintf("%03d", code)
	for _, cur := range Currencies {
		if cur.GetIsoNumeric() == iso {
			return cur, true
		}
	}
	return nil, false
}

// GetBySymbol returns all currencies that use the given symbol
func GetBySymbol(symbol string) []Currency {
	var result []Currency
	for _, cur := range Currencies {
		if cur.GetSymbol() == symbol {
			result = append(result, cur)
		}
	}
	return result
}

// Exists returns true if the given currency code is supported, case-insensitive
func Exists(code string) bool {
	_, ok := GetByCode(code)
	return ok
}

// AllCodes returns a sorted slice of all currency codes
func AllCodes() []string {
	codes := make([]string, 0, len(Currencies))
	for code := range Currencies {
		codes = append(codes, code)
	}

	// Sort the codes
	for i := 0; i < len(codes)-1; i++ {
		for j := i + 1; j < len(codes); j++ {
			if codes[i] > codes[j] {
				codes[i], codes[j] = codes[j], codes[i]
			}
		}
	}

	return codes
}
