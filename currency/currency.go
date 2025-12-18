package currency

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
)

//go:embed data/currency.json
var currenciesJSON []byte

// rawModel matches the structure used in the PHP build file: { "Names": { "USD": [symbol, name], ... } }
type rawModel struct {
	Names map[string][2]string `json:"Names"`
}

// Currency represents basic currency metadata used by this package.
type Currency struct {
	Code   string
	Symbol string
	Name   string
}

var (
	data       rawModel
	codeToInfo map[string]Currency
)

func init() {
	_ = json.Unmarshal(currenciesJSON, &data)
	codeToInfo = make(map[string]Currency, len(data.Names))
	for code, arr := range data.Names {
		c := Currency{
			Code:   code,
			Symbol: arr[0],
			Name:   arr[1],
		}
		codeToInfo[c.Code] = c
	}
}

// Get returns a Currency by ISO code (case-insensitive).
func Get(code string) (Currency, error) {
	if code == "" {
		return Currency{}, errors.New("currency code is required")
	}
	c, ok := codeToInfo[strings.ToUpper(code)]
	if !ok {
		return Currency{}, errors.New("unsupported currency: " + code)
	}
	return c, nil
}

// ListAll returns a sorted list of all currency codes to names.
func ListAll() map[string]string {
	// keep map output stable by sorting keys first
	keys := make([]string, 0, len(codeToInfo))
	for k := range codeToInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make(map[string]string, len(keys))
	for _, k := range keys {
		out[k] = codeToInfo[k].Name
	}
	return out
}

// Common returns a subset of common currencies similar to PHP CurrencyHelper::commonCurrencies().
func Common() map[string]string {
	codes := []string{"USD", "GBP", "EUR", "AUD", "CAD"}
	out := make(map[string]string, len(codes))
	for _, c := range codes {
		if cur, ok := codeToInfo[c]; ok {
			out[c] = cur.Name
		}
	}
	return out
}

// Format returns a string formatted amount using the currency's symbol and code.
// Behavior mirrors the defaults of the PHP implementation for most currencies:
// - uses 2 decimals, '.' decimal separator, ',' thousands separator
// - negative values use a leading '-' for most currencies
// - USD specifically wraps negatives in parentheses: ($1.00)
// You can choose to display symbol and/or code.
func Format(c Currency, amount float64, showSymbol, showCode bool) string {
	abs := amount
	if abs < 0 {
		abs = -abs
	}
	num := numberFormat(abs, 2, '.', ',')
	symbol := ""
	if showSymbol {
		symbol = c.Symbol
	}
	code := ""
	if showCode {
		code = c.Code
	}
	// base format: optional '-', then {symbol}{number} and optional space + {code}
	core := symbol + num
	if showCode && code != "" {
		core += " " + code
	}

	if amount < 0 {
		if c.Code == "USD" {
			return "(" + core + ")"
		}
		return "-" + core
	}
	return core
}

// numberFormat replicates a simple 1000s format with fixed decimals.
func numberFormat(n float64, decimals int, dec byte, thou byte) string {
	// Round to required decimals using sprintf-like behavior.
	// Implement minimal formatter without importing fmt for performance and stability.
	// We'll use fmt for simplicity as minimalism is not critical here.
	// NOTE: Keeping implementation simple and dependable.
	return simpleNumberFormat(n, decimals, dec, thou)
}

// simpleNumberFormat uses fmt and string post-processing for thousand separators.
func simpleNumberFormat(n float64, decimals int, dec byte, thou byte) string {
	// Use fmt to get a fixed decimal string, then inject thousands separators.
	// We avoid locale complexity.
	s := formatFixed(n, decimals)
	// split integer and fraction
	i := strings.IndexByte(s, '.')
	intPart := s
	frac := ""
	if i >= 0 {
		intPart = s[:i]
		frac = s[i+1:]
	}
	// insert thousands separators into intPart
	var b strings.Builder
	l := len(intPart)
	for idx, r := range intPart {
		// r is rune but input is ASCII digits, keep it simple
		b.WriteRune(r)
		pos := idx + 1
		if l > 3 && pos < l && (l-pos)%3 == 0 {
			b.WriteByte(thou)
		}
	}
	if decimals > 0 {
		b.WriteByte(dec)
		// pad or trim fraction to required length
		if len(frac) < decimals {
			b.WriteString(frac)
			for i := len(frac); i < decimals; i++ {
				b.WriteByte('0')
			}
		} else {
			b.WriteString(frac[:decimals])
		}
	}
	return b.String()
}

func formatFixed(n float64, decimals int) string {
	// Use fmt to get a fixed decimal string
	return fmt.Sprintf("%.*f", decimals, n)
}
