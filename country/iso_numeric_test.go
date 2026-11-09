package country

import (
	"strconv"
	"testing"
)

func TestGetIsoNumeric(t *testing.T) {
	for code, want := range map[string]string{"BE": "056", "AT": "040", "AD": "020", "US": "840", "GB": "826", "XK": ""} {
		if got := List[code].GetIsoNumeric(); got != want {
			t.Errorf("%s GetIsoNumeric() = %q, want %q", code, got, want)
		}
	}
}

// Every numeric code is three digits that Lookup resolves back to its country.
func TestIsoNumericRoundTrip(t *testing.T) {
	for code, c := range List {
		n := c.GetIsoNumeric()
		if n == "" {
			continue
		}
		if _, err := strconv.Atoi(n); err != nil || len(n) != 3 {
			t.Errorf("%s GetIsoNumeric() = %q, want three digits", code, n)
		}
		if got, ok := Lookup(n); !ok || got.GetIso2() != code {
			t.Errorf("Lookup(%q) did not find %s", n, code)
		}
	}
}
