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
		{"GB", "Hampshire", "GB-HAM"},
		{"GB", "Buckinghamshire", "GB-BKM"},
		{"gb", "ham", "GB-HAM"},
		{"GB", "GB-ENG", "GB-ENG"},
		{"FR", "Île-de-France", "FR-IDF"},
		{"FR", "ile de france", "FR-IDF"},
		{"EE", "661", "EE-661"},
	} {
		s, ok := Lookup(tt.country, tt.in)
		if !ok || s.Code != tt.want {
			t.Errorf("Lookup(%q, %q) = %v, %v; want %s", tt.country, tt.in, s, ok, tt.want)
		}
	}
	for _, tt := range []struct{ country, in string }{
		{"US", ""}, {"US", "Ontario"}, {"CA", "FL"}, {"US", "Buckinghamshire"}, {"ZZ", "FL"},
		// Two Estonian subdivisions are both named Rakvere, so the name finds neither.
		{"EE", "Rakvere"},
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
	for country, want := range map[string]int{"US": 57, "ca": 13, "GB": 221} {
		if n := len(List(country)); n != want {
			t.Errorf("%s has %d subdivisions, want %d", country, n, want)
		}
	}
	if n := len(List("ZZ")); n != 0 {
		t.Errorf("ZZ has %d subdivisions, want 0", n)
	}
}

// Every subdivision is found by its own code, and no name hides behind a code.
func TestEveryCodeFindsItself(t *testing.T) {
	loadOnce.Do(load)
	for country, idx := range byCountry {
		for _, s := range idx.list {
			if got, ok := Lookup(country, s.Code); !ok || got.Code != s.Code {
				t.Errorf("Lookup(%s, %s) = %v, %v", country, s.Code, got, ok)
			}
		}
	}
}
