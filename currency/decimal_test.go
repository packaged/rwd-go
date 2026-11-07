package currency

import (
	"errors"
	"math"
	"testing"
)

func mustCurrency(t *testing.T, code string) Currency {
	t.Helper()
	c, ok := GetByCode(code)
	if !ok {
		t.Fatalf("currency %s not found", code)
	}
	return c
}

func TestFormatDecimal(t *testing.T) {
	tests := []struct {
		code   string
		amount int64
		want   string
	}{
		{"USD", 2999, "29.99"},
		{"USD", 4000, "40.00"},
		{"USD", 5, "0.05"},
		{"USD", 0, "0.00"},
		{"USD", -5, "-0.05"},
		{"USD", -2999, "-29.99"},
		{"USD", 123456789, "1234567.89"},
		{"JPY", 1000, "1000"},
		{"JPY", -7, "-7"},
		{"BHD", 1234, "1.234"},
		{"BHD", 5, "0.005"},
		{"USD", math.MinInt64, "-92233720368547758.08"},
		{"USD", math.MaxInt64, "92233720368547758.07"},
	}
	for _, tt := range tests {
		if got := mustCurrency(t, tt.code).FormatDecimal(tt.amount); got != tt.want {
			t.Errorf("%s FormatDecimal(%d) = %q, want %q", tt.code, tt.amount, got, tt.want)
		}
	}
}

func TestParseDecimal(t *testing.T) {
	tests := []struct {
		code    string
		in      string
		want    int64
		wantErr error
	}{
		{"USD", "29.99", 2999, nil},
		{"USD", "29.9", 2990, nil},
		{"USD", "29", 2900, nil},
		{"USD", "29.", 2900, nil},
		{"USD", ".5", 50, nil},
		{"USD", "0.05", 5, nil},
		{"USD", "-0.05", -5, nil},
		{"USD", "+1.00", 100, nil},
		{"USD", " 1.00 ", 100, nil},
		{"USD", "29.9900", 2999, nil},
		{"USD", "29.999", 0, ErrDecimalPrecision},
		{"JPY", "1000", 1000, nil},
		{"JPY", "1000.0", 1000, nil},
		{"JPY", "1000.5", 0, ErrDecimalPrecision},
		{"BHD", "1.234", 1234, nil},
		{"USD", "", 0, ErrInvalidDecimal},
		{"USD", "-", 0, ErrInvalidDecimal},
		{"USD", ".", 0, ErrInvalidDecimal},
		{"USD", "1,000.00", 0, ErrInvalidDecimal},
		{"USD", "1e3", 0, ErrInvalidDecimal},
		{"USD", "1.2.3", 0, ErrInvalidDecimal},
		{"USD", "--1", 0, ErrInvalidDecimal},
		{"USD", "92233720368547758.07", math.MaxInt64, nil},
		{"USD", "92233720368547758.08", 0, ErrDecimalOverflow},
	}
	for _, tt := range tests {
		got, err := mustCurrency(t, tt.code).ParseDecimal(tt.in)
		if !errors.Is(err, tt.wantErr) || got != tt.want {
			t.Errorf("%s ParseDecimal(%q) = %d, %v; want %d, %v", tt.code, tt.in, got, err, tt.want, tt.wantErr)
		}
	}
}

func TestDecimalRoundTrip(t *testing.T) {
	for _, code := range AllCodes() {
		c := mustCurrency(t, code)
		for _, amount := range []int64{0, 1, -1, 5, 99, 100, 2999, -123456} {
			got, err := c.ParseDecimal(c.FormatDecimal(amount))
			if err != nil || got != amount {
				t.Errorf("%s round trip of %d = %d, %v", code, amount, got, err)
			}
		}
	}
}

func TestGetByNumericCode(t *testing.T) {
	for code, want := range map[uint16]string{840: "USD", 978: "EUR", 392: "JPY", 826: "GBP"} {
		c, ok := GetByNumericCode(code)
		if !ok || c.GetCode() != want {
			t.Errorf("GetByNumericCode(%d) = %v, %v; want %s", code, c, ok, want)
		}
	}
	if _, ok := GetByNumericCode(0); ok {
		t.Error("GetByNumericCode(0) found a currency")
	}
}

func TestGetByCodeIgnoresCase(t *testing.T) {
	for _, in := range []string{"usd", "Usd", " USD "} {
		if c, ok := GetByCode(in); !ok || c.GetCode() != "USD" {
			t.Errorf("GetByCode(%q) = %v, %v; want USD", in, c, ok)
		}
	}
}
