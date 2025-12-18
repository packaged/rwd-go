package currency

import "testing"

func TestGet(t *testing.T) {
	c, err := Get("usd")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Code != "USD" || c.Symbol != "$" || c.Name != "US Dollar" {
		t.Fatalf("unexpected currency: %+v", c)
	}
}

func TestCommon(t *testing.T) {
	m := Common()
	want := []string{"USD", "GBP", "EUR", "AUD", "CAD"}
	for _, k := range want {
		if _, ok := m[k]; !ok {
			t.Fatalf("missing common currency %s", k)
		}
	}
}

func TestFormat(t *testing.T) {
	usd, _ := Get("USD")
	s := Format(usd, 1234.5, true, false)
	if s != "$1,234.50" {
		t.Fatalf("unexpected format: %s", s)
	}
	s = Format(usd, -1234.5, true, true)
	if s != "($1,234.50 USD)" {
		t.Fatalf("unexpected negative USD format: %s", s)
	}

	eur, _ := Get("EUR")
	s = Format(eur, -9876.5, true, true)
	if s != "-€9,876.50 EUR" {
		t.Fatalf("unexpected negative EUR format: %s", s)
	}
}
