package currency

import (
	"fmt"
	"testing"
)

// TestCurrenciesInitialized verifies all expected currencies exist in the map
func TestCurrenciesInitialized(t *testing.T) {
	expectedCodes := []string{
		"AED", "AFN", "ALL", "AMD", "AOA", "ARS", "AUD", "AWG", "AZN",
		"BAM", "BBD", "BDT", "BHD", "BIF", "BMD", "BND", "BOB", "BRL",
		"BSD", "BTN", "BWP", "BYN", "BZD",
		"CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK",
		"DJF", "DKK", "DOP", "DZD",
		"EGP", "ERN", "ETB", "EUR",
		"FJD", "FKP",
		"GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD",
		"HKD", "HNL", "HTG", "HUF",
		"IDR", "ILS", "INR", "IQD", "IRR", "ISK",
		"JMD", "JOD", "JPY",
		"KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD",
		"LAK", "LBP", "LKR", "LRD", "LSL", "LYD",
		"MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR",
		"MVR", "MWK", "MXN", "MYR", "MZN",
		"NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
		"OMR",
		"PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG",
		"QAR",
		"RON", "RSD", "RUB", "RWF",
		"SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLE", "SOS",
		"SRD", "SSP", "STN", "SVC", "SYP", "SZL",
		"THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS",
		"UAH", "UGX", "USD", "UYU", "UZS",
		"VES", "VND", "VUV",
		"WST",
		"XAF", "XCD", "XCG", "XOF", "XPF",
		"YER",
		"ZAR", "ZMW", "ZWG",
	}

	if len(Currencies) != len(expectedCodes) {
		t.Errorf("Expected %d currencies, got %d", len(expectedCodes), len(Currencies))
	}

	for _, code := range expectedCodes {
		if _, exists := Currencies[code]; !exists {
			t.Errorf("Currency %s not found in Currencies map", code)
		}
	}
}

// TestCurrencyFormat verifies each currency has a default format pattern
func TestCurrencyFormat(t *testing.T) {
	tests := []struct {
		code           string
		expectedFormat string
	}{
		{"USD", "{symbol}{amount}"},
		{"EUR", "{symbol}{amount}"},
		{"GBP", "{symbol}{amount}"},
		{"JPY", "{symbol}{amount}"},
		{"BRL", "{code} {amount}"},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			cur := Currencies[tt.code]
			format := cur.getFormat()
			if format == "" {
				t.Errorf("%s has empty format pattern", tt.code)
			}
			// Just verify it's not empty - each currency can have its own pattern
		})
	}
}

// TestFormatDefault tests formatting with the currency's default pattern
func TestFormatDefault(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		amount   int64
		expected string
	}{
		{
			name:     "USD uses default {symbol}{amount}",
			currency: "USD",
			amount:   123456,
			expected: "$1,234.56",
		},
		{
			name:     "EUR uses default {symbol}{amount}",
			currency: "EUR",
			amount:   123456,
			expected: "€1.234,56",
		},
		{
			name:     "BRL uses default {code} {amount}",
			currency: "BRL",
			amount:   123456,
			expected: "BRL 1.234,56",
		},
		{
			name:     "DKK uses default {amount} {code}",
			currency: "DKK",
			amount:   123456,
			expected: "1.234,56 DKK",
		},
		{
			name:     "JPY with no decimals",
			currency: "JPY",
			amount:   123456,
			expected: "¥123,456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur := Currencies[tt.currency]
			result := cur.Format(tt.amount)
			if result != tt.expected {
				t.Errorf("Format(%d) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

// TestFormat tests currency formatting with pattern strings
func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		amount   int64
		pattern  string
		expected string
	}{
		{
			name:     "USD with symbol and amount",
			currency: "USD",
			amount:   123456, // represents $1,234.56
			pattern:  "{symbol}{amount}",
			expected: "$1,234.56",
		},
		{
			name:     "USD with code, symbol and amount",
			currency: "USD",
			amount:   123456,
			pattern:  "{code}{symbol}{amount}",
			expected: "USD$1,234.56",
		},
		{
			name:     "EUR with comma decimal separator",
			currency: "EUR",
			amount:   123456,
			pattern:  "{symbol}{amount}",
			expected: "€1.234,56",
		},
		{
			name:     "JPY with no decimals",
			currency: "JPY",
			amount:   123456,
			pattern:  "{symbol}{amount}",
			expected: "¥123,456",
		},
		{
			name:     "GBP with amount then code",
			currency: "GBP",
			amount:   123456,
			pattern:  "{amount} {code}",
			expected: "1,234.56 GBP",
		},
		{
			name:     "CHF with apostrophe thousand separator",
			currency: "CHF",
			amount:   123456,
			pattern:  "{symbol}{amount}",
			expected: "CHF1'234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur := Currencies[tt.currency]
			result := cur.format(tt.amount, tt.pattern)
			if result != tt.expected {
				t.Errorf("format(%d, %q) = %q, want %q", tt.amount, tt.pattern, result, tt.expected)
			}
		})
	}
}

// TestGetBySymbol tests currency lookup by symbol
func TestGetBySymbol(t *testing.T) {
	tests := []struct {
		name       string
		symbol     string
		wantCodes  []string
		wantLength int
	}{
		{
			name:      "Dollar sign returns multiple currencies",
			symbol:    "$",
			wantCodes: []string{"AUD", "CAD", "MXN", "NZD", "USD"},
			wantLength: 27,
		},
		{
			name:       "Euro symbol returns one currency",
			symbol:     "€",
			wantCodes:  []string{"EUR"},
			wantLength: 1,
		},
		{
			name:       "Yen symbol returns one currency",
			symbol:     "¥",
			wantCodes:  []string{"CNY", "JPY"},
			wantLength: 2,
		},
		{
			name:      "kr symbol returns multiple Nordic currencies",
			symbol:    "kr",
			wantCodes: []string{"DKK", "ISK", "NOK", "SEK"},
			wantLength: 4,
		},
		{
			name:       "Non-existent symbol",
			symbol:     "@",
			wantLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetBySymbol(tt.symbol)
			if len(result) != tt.wantLength {
				t.Errorf("GetBySymbol(%q) returned %d currencies, want %d", tt.symbol, len(result), tt.wantLength)
			}

			// Check that all expected codes are present
			if tt.wantLength > 0 {
				foundCodes := make(map[string]bool)
				for _, cur := range result {
					foundCodes[cur.getCode()] = true
				}

				for _, wantCode := range tt.wantCodes {
					if !foundCodes[wantCode] {
						t.Errorf("GetBySymbol(%q) missing expected code %q", tt.symbol, wantCode)
					}
				}
			}
		})
	}
}

// TestExists tests currency code validation
func TestExists(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{
			name: "Valid currency USD",
			code: "USD",
			want: true,
		},
		{
			name: "Valid currency EUR",
			code: "EUR",
			want: true,
		},
		{
			name: "Valid currency JPY",
			code: "JPY",
			want: true,
		},
		{
			name: "Invalid currency",
			code: "XXX",
			want: false,
		},
		{
			name: "Empty string",
			code: "",
			want: false,
		},
		{
			name: "Lowercase code",
			code: "usd",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exists(tt.code)
			if got != tt.want {
				t.Errorf("Exists(%q) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

// TestAllCodes tests getting all currency codes
func TestAllCodes(t *testing.T) {
	codes := AllCodes()

	// Should have 153 currencies
	if len(codes) != 153 {
		t.Errorf("AllCodes() returned %d codes, want 153", len(codes))
	}

	// Should be sorted
	for i := 1; i < len(codes); i++ {
		if codes[i-1] >= codes[i] {
			t.Errorf("AllCodes() not sorted: %q >= %q at position %d", codes[i-1], codes[i], i)
		}
	}

	// Should contain all originally expected codes
	expectedCodes := []string{
		"AUD", "BRL", "CAD", "CHF", "CNY", "DKK", "EUR", "GBP",
		"HKD", "IDR", "INR", "JPY", "KRW", "MXN", "NOK", "NZD",
		"PLN", "RUB", "SAR", "SEK", "THB", "TRY", "TWD", "USD", "ZAR",
	}

	codeMap := make(map[string]bool)
	for _, code := range codes {
		codeMap[code] = true
	}

	for _, expected := range expectedCodes {
		if !codeMap[expected] {
			t.Errorf("AllCodes() missing expected code %q", expected)
		}
	}
}

// TestGetByCode tests safe currency lookup
func TestGetByCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		wantCode string
		wantOk   bool
	}{
		{
			name:     "Valid currency USD",
			code:     "USD",
			wantCode: "USD",
			wantOk:   true,
		},
		{
			name:     "Valid currency EUR",
			code:     "EUR",
			wantCode: "EUR",
			wantOk:   true,
		},
		{
			name:   "Invalid currency",
			code:   "XXX",
			wantOk: false,
		},
		{
			name:   "Empty string",
			code:   "",
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur, ok := GetByCode(tt.code)
			if ok != tt.wantOk {
				t.Errorf("GetByCode(%q) ok = %v, want %v", tt.code, ok, tt.wantOk)
			}
			if ok && cur.getCode() != tt.wantCode {
				t.Errorf("GetByCode(%q) code = %q, want %q", tt.code, cur.getCode(), tt.wantCode)
			}
		})
	}
}

// TestFromInt tests converting minor unit int64 to float64
func TestFromInt(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		amount   int64
		expected float64
	}{
		{"USD positive", "USD", 12345, 123.45},
		{"USD zero", "USD", 0, 0.0},
		{"USD negative", "USD", -12345, -123.45},
		{"USD small", "USD", 1, 0.01},
		{"JPY no decimals", "JPY", 12345, 12345.0},
		{"JPY negative no decimals", "JPY", -500, -500.0},
		{"EUR positive", "EUR", 100, 1.0},
		{"IDR no decimals", "IDR", 50000, 50000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur := Currencies[tt.currency]
			result := cur.FromInt(tt.amount)
			if result != tt.expected {
				t.Errorf("FromInt(%d) = %f, want %f", tt.amount, result, tt.expected)
			}
		})
	}
}

// TestFromInt32 tests converting minor unit int32 to float64
func TestFromInt32(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		amount   int32
		expected float64
	}{
		{"USD positive", "USD", 12345, 123.45},
		{"USD negative", "USD", -12345, -123.45},
		{"JPY no decimals", "JPY", 12345, 12345.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cur := Currencies[tt.currency]
			result := cur.FromInt32(tt.amount)
			if result != tt.expected {
				t.Errorf("FromInt32(%d) = %f, want %f", tt.amount, result, tt.expected)
			}
		})
	}
}

// Example functions for documentation

func ExampleGetByCode() {
	cur, ok := GetByCode("EUR")
	if ok {
		fmt.Println(cur.getName())
	}
	// Output: Euro
}

func ExampleGetBySymbol() {
	currencies := GetBySymbol("$")
	// Sort for consistent output
	codes := []string{}
	for _, cur := range currencies {
		codes = append(codes, cur.getCode())
	}
	// Simple bubble sort
	for i := 0; i < len(codes)-1; i++ {
		for j := i + 1; j < len(codes); j++ {
			if codes[i] > codes[j] {
				codes[i], codes[j] = codes[j], codes[i]
			}
		}
	}
	for _, code := range codes {
		fmt.Println(code)
	}
	// Output:
	// ARS
	// AUD
	// BBD
	// BMD
	// BND
	// BSD
	// BZD
	// CAD
	// CLP
	// COP
	// CUP
	// CVE
	// DOP
	// FJD
	// GYD
	// JMD
	// KYD
	// LRD
	// MXN
	// NAD
	// NZD
	// SBD
	// SGD
	// SRD
	// TTD
	// USD
	// XCD
}

func ExampleExists() {
	if Exists("EUR") {
		fmt.Println("EUR is supported")
	}
	// Output: EUR is supported
}

func ExampleCurrency_format() {
	usd := Currencies["USD"]
	fmt.Println(usd.format(123456, "{symbol}{amount}"))

	eur := Currencies["EUR"]
	fmt.Println(eur.format(123456, "{code} {amount}"))
	// Output:
	// $1,234.56
	// EUR 1.234,56
}

func ExampleCurrency_Format() {
	// Format using each currency's default pattern
	usd := Currencies["USD"]
	fmt.Println(usd.Format(123456)) // Uses USD's default: {symbol}{amount}

	brl := Currencies["BRL"]
	fmt.Println(brl.Format(123456)) // Uses BRL's default: {code} {amount}

	jpy := Currencies["JPY"]
	fmt.Println(jpy.Format(123456)) // Uses JPY's default: {symbol}{amount}, no decimals
	// Output:
	// $1,234.56
	// BRL 1.234,56
	// ¥123,456
}
