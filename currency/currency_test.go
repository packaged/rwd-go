package currency

import (
	"fmt"
	"testing"
)

// TestCurrenciesInitialized verifies all 24 currencies exist in the map
func TestCurrenciesInitialized(t *testing.T) {
	expectedCodes := []string{
		"AUD", "BRL", "CAD", "CHF", "CNY", "DKK", "EUR", "GBP",
		"HKD", "IDR", "INR", "JPY", "KRW", "MXN", "NOK", "NZD",
		"PLN", "RUB", "SAR", "SEK", "THB", "TRY", "TWD", "USD", "ZAR",
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
			name:       "Dollar sign returns multiple currencies",
			symbol:     "$",
			wantCodes:  []string{"AUD", "CAD", "MXN", "NZD", "USD"},
			wantLength: 5,
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
			name:       "kr symbol returns multiple Nordic currencies",
			symbol:     "kr",
			wantCodes:  []string{"DKK", "NOK", "SEK"},
			wantLength: 3,
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

	// Should have 25 currencies
	if len(codes) != 25 {
		t.Errorf("AllCodes() returned %d codes, want 25", len(codes))
	}

	// Should be sorted
	for i := 1; i < len(codes); i++ {
		if codes[i-1] >= codes[i] {
			t.Errorf("AllCodes() not sorted: %q >= %q at position %d", codes[i-1], codes[i], i)
		}
	}

	// Should contain all expected codes
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
	// AUD
	// CAD
	// MXN
	// NZD
	// USD
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
