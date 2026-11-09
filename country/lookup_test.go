package country

import "testing"

func TestGetByIso3(t *testing.T) {
	for in, want := range map[string]string{"GBR": "GB", "gbr": "GB", " usa ": "US", "XKX": "XK"} {
		if c, ok := GetByIso3(in); !ok || c.GetIso2() != want {
			t.Errorf("GetByIso3(%q) = %v, %v; want %s", in, c, ok, want)
		}
	}
	for _, in := range []string{"", "GB", "ZZZ"} {
		if _, ok := GetByIso3(in); ok {
			t.Errorf("GetByIso3(%q) found a country", in)
		}
	}
}

func TestGetByNumericCode(t *testing.T) {
	for in, want := range map[uint16]string{826: "GB", 840: "US", 56: "BE", 70: "BA", 4: "AF"} {
		if c, ok := GetByNumericCode(in); !ok || c.GetIso2() != want {
			t.Errorf("GetByNumericCode(%d) = %v, %v; want %s", in, c, ok, want)
		}
	}
	if c, ok := GetByNumericCode(0); ok {
		t.Errorf("GetByNumericCode(0) = %s, want not found", c.GetIso2())
	}
}

func TestGetByName(t *testing.T) {
	for in, want := range map[string]string{
		"United Kingdom": "GB", "united states": "US", " Germany ": "DE",
		"Czechia": "CZ", "Czech Republic": "CZ", "Ivory Coast": "CI", "Côte d’Ivoire": "CI", "Côte d'Ivoire": "CI",
		"Swaziland": "SZ", "Eswatini": "SZ", "Russian Federation": "RU", "Antigua and Barbuda": "AG",
	} {
		if c, ok := GetByName(in); !ok || c.GetIso2() != want {
			t.Errorf("GetByName(%q) = %v, %v; want %s", in, c, ok, want)
		}
	}
	for _, in := range []string{"Atlantis", "", "  "} {
		if c, ok := GetByName(in); ok {
			t.Errorf("GetByName(%q) = %v; want no country", in, c)
		}
	}
}

func TestLookup(t *testing.T) {
	tests := map[string]string{
		"GB": "GB", "gb": "GB", "UK": "GB", "GBR": "GB", "826": "GB", "056": "BE", " usa ": "US", "XK": "XK",
	}
	for in, want := range tests {
		if c, ok := Lookup(in); !ok || c.GetIso2() != want {
			t.Errorf("Lookup(%q) = %v, %v; want %s", in, c, ok, want)
		}
	}
	for _, in := range []string{"", "0", "999", "70000", "G", "GBRX", "ZZ", "ZZZ"} {
		if c, ok := Lookup(in); ok {
			t.Errorf("Lookup(%q) = %s, want not found", in, c.GetIso2())
		}
	}
}

// Every alpha-3 and every non-zero numeric code identifies exactly one country, so the
// lookups by them are unambiguous.
func TestCodesAreUnique(t *testing.T) {
	iso3 := map[string]string{}
	numeric := map[string]string{}
	for iso2, c := range List {
		if prev, dup := iso3[c.GetIso3()]; dup {
			t.Errorf("alpha-3 %s is used by %s and %s", c.GetIso3(), prev, iso2)
		}
		iso3[c.GetIso3()] = iso2
		if n := c.GetIsoNumeric(); n != "" {
			if prev, dup := numeric[n]; dup {
				t.Errorf("numeric %s is used by %s and %s", n, prev, iso2)
			}
			numeric[n] = iso2
		}
	}
}
