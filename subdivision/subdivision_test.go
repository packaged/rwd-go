package subdivision

import "testing"

func TestLookup(t *testing.T) {
	for _, tt := range []struct{ country, in, want string }{
		{"US", "FL", "US-FL"},
		{"us", "fl", "US-FL"},
		{"US", "US-FL", "US-FL"},
		{"US", "Florida", "US-FL"},
		{"US", " florida ", "US-FL"},
		{"US", "New York", "US-NY"},
		{"US", "newyork", "US-NY"},
		{"US", "District of Columbia", "US-DC"},
		{"US", "Washington D.C.", "US-DC"},
		{"US", "Washington", "US-WA"},
		{"US", "Puerto Rico", "US-PR"},
		{"CA", "ON", "CA-ON"},
		{"CA", "Québec", "CA-QC"},
		{"CA", "Quebec", "CA-QC"},
		{"CA", "QC", "CA-QC"},
		{"CA", "Newfoundland", "CA-NL"},
		{"CA", "Île-du-Prince-Édouard", "CA-PE"},
		{"CA", "british columbia", "CA-BC"},
	} {
		s, ok := Lookup(tt.country, tt.in)
		if !ok || s.Code != tt.want {
			t.Errorf("Lookup(%q, %q) = %v, %v; want %s", tt.country, tt.in, s, ok, tt.want)
		}
	}
	for _, tt := range []struct{ country, in string }{
		{"US", ""}, {"US", "Ontario"}, {"CA", "FL"}, {"GB", "Hampshire"}, {"US", "Buckinghamshire"},
	} {
		if s, ok := Lookup(tt.country, tt.in); ok {
			t.Errorf("Lookup(%q, %q) = %s, want none", tt.country, tt.in, s.Code)
		}
	}
}

func TestCodes(t *testing.T) {
	s, _ := Lookup("US", "FL")
	if s.Country() != "US" || s.LocalCode() != "FL" || s.Name != "Florida" {
		t.Errorf("FL = %s %s %s", s.Country(), s.LocalCode(), s.Name)
	}
}

func TestList(t *testing.T) {
	if n := len(List("US")); n != 57 {
		t.Errorf("US has %d subdivisions, want 57 (50 states, DC, 6 outlying areas)", n)
	}
	if n := len(List("ca")); n != 13 {
		t.Errorf("CA has %d subdivisions, want 13", n)
	}
	if n := len(List("GB")); n != 0 {
		t.Errorf("GB has %d subdivisions, want 0 (no data yet)", n)
	}
}

// No two subdivisions of a country share a code, name or alias once folded, so a lookup
// never has two answers.
func TestKeysAreUnique(t *testing.T) {
	for country, list := range byCountry {
		seen := map[string]string{}
		for _, s := range list {
			keys := append([]string{s.LocalCode(), s.Name}, s.aliases...)
			for _, k := range keys {
				if prev, dup := seen[fold(k)]; dup && prev != s.Code {
					t.Errorf("%s: %q is both %s and %s", country, k, prev, s.Code)
				}
				seen[fold(k)] = s.Code
			}
		}
	}
}
