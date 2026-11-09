package currency

import "testing"

func TestGetIsoNumeric(t *testing.T) {
	for code, want := range map[string]string{"ALL": "008", "AUD": "036", "BHD": "048", "USD": "840", "EUR": "978"} {
		if got := Currencies[code].GetIsoNumeric(); got != want {
			t.Errorf("%s GetIsoNumeric() = %q, want %q", code, got, want)
		}
	}
	for code, cur := range Currencies {
		if n := cur.GetIsoNumeric(); len(n) != 3 {
			t.Errorf("%s GetIsoNumeric() = %q, want three digits", code, n)
		}
	}
}
