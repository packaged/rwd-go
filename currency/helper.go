package currency

// GetByCode returns the currency for the given code and a boolean indicating if found
func GetByCode(code string) (Currency, bool) {
	cur, ok := Currencies[code]
	return cur, ok
}

// GetBySymbol returns all currencies that use the given symbol
func GetBySymbol(symbol string) []Currency {
	var result []Currency
	for _, cur := range Currencies {
		if cur.getSymbol() == symbol {
			result = append(result, cur)
		}
	}
	return result
}

// Exists returns true if the given currency code is supported
func Exists(code string) bool {
	_, ok := Currencies[code]
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
