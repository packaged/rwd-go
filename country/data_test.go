package country

import (
	"slices"
	"testing"

	"github.com/packaged/rwd-go/currency"
)

// Every country's currency is one the currency package holds.
func TestCurrenciesResolve(t *testing.T) {
	for code, c := range List {
		if cur := c.GetCurrencyCode(); cur != "" {
			if _, ok := currency.GetByCode(cur); !ok {
				t.Errorf("%s: currency %s is not in the currency package", code, cur)
			}
		}
	}
}

func TestEUCountryCodes(t *testing.T) {
	if len(EUCountryCodes) != 27 {
		t.Errorf("EU has %d members, want 27", len(EUCountryCodes))
	}
	for _, code := range []string{"GB", "XK"} {
		if slices.Contains(EUCountryCodes, code) {
			t.Errorf("EUCountryCodes holds %s", code)
		}
	}
}

// DialCodes holds every country with a dialling code under that code.
func TestDialCodes(t *testing.T) {
	for code, c := range List {
		if p := c.GetDialPrefix(); p != 0 && !slices.Contains(DialCodes[p], code) {
			t.Errorf("DialCodes[%d] is missing %s", p, code)
		}
	}
	if got := DialCodes[44]; !slices.Equal(got, []string{"GB", "GG", "IM", "JE"}) {
		t.Errorf("DialCodes[44] = %v", got)
	}
}
