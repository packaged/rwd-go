// Package subdivision holds ISO 3166-2 country subdivisions (states, provinces, counties
// and the like) and finds them by code or name. The data comes from Unicode CLDR.
package subdivision

//go:generate sh -c "go run ./internal/cldrgen -release release-49 > subdivisions.tsv"

import (
	_ "embed"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

//go:embed subdivisions.tsv
var subdivisionsTSV string

// Subdivision is an ISO 3166-2 country subdivision.
type Subdivision struct {
	// Code is the full ISO 3166-2 code, such as "US-FL".
	Code string
	// Name is the English name, such as "Florida".
	Name string
}

// Country returns the ISO 3166-1 alpha-2 code of the subdivision's country: "US" for
// "US-FL".
func (s Subdivision) Country() string {
	country, _, _ := strings.Cut(s.Code, "-")
	return country
}

// LocalCode returns the code within its country: "FL" for "US-FL".
func (s Subdivision) LocalCode() string {
	_, local, _ := strings.Cut(s.Code, "-")
	return local
}

type countryIndex struct {
	list   []Subdivision
	byCode map[string]int
	// byName maps a folded name or alias to its subdivision, or to -1 when the name
	// belongs to more than one.
	byName map[string]int
}

var (
	loadOnce  sync.Once
	byCountry map[string]*countryIndex
)

func load() {
	byCountry = map[string]*countryIndex{}
	add := func(code, name string) {
		country, local, ok := strings.Cut(code, "-")
		if !ok {
			return
		}
		idx := byCountry[country]
		if idx == nil {
			idx = &countryIndex{byCode: map[string]int{}, byName: map[string]int{}}
			byCountry[country] = idx
		}
		if _, dup := idx.byCode[fold(local)]; dup {
			return
		}
		idx.byCode[fold(local)] = len(idx.list)
		idx.byCode[fold(code)] = len(idx.list)
		idx.list = append(idx.list, Subdivision{Code: code, Name: name})
	}
	for _, line := range strings.Split(subdivisionsTSV, "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		code, name, _ := strings.Cut(line, "\t")
		add(code, name)
	}
	for _, s := range extra {
		add(s.Code, s.Name)
	}
	for _, idx := range byCountry {
		for i, s := range idx.list {
			nameKey(idx, s.Name, i)
		}
	}
	for code, names := range aliases {
		country, _, _ := strings.Cut(code, "-")
		if idx := byCountry[country]; idx != nil {
			if i, ok := idx.byCode[fold(code)]; ok {
				for _, n := range names {
					nameKey(idx, n, i)
				}
			}
		}
	}
}

func nameKey(idx *countryIndex, name string, i int) {
	key := fold(name)
	if key == "" {
		return
	}
	if prev, seen := idx.byName[key]; seen && prev != i {
		idx.byName[key] = -1
		return
	}
	idx.byName[key] = i
}

// List returns the subdivisions of a country, by its ISO 3166-1 alpha-2 code,
// case-insensitive.
func List(country string) []Subdivision {
	loadOnce.Do(load)
	if idx := byCountry[strings.ToUpper(strings.TrimSpace(country))]; idx != nil {
		return append([]Subdivision(nil), idx.list...)
	}
	return nil
}

// Lookup finds a country's subdivision by its local code ("FL"), full code ("US-FL") or
// name ("Florida"), ignoring case, accents, spacing and punctuation, so "Québec",
// "quebec" and "QC" all find Quebec. A code is matched before a name, and a name that
// more than one of the country's subdivisions share finds none. The country is an
// ISO 3166-1 alpha-2 code, case-insensitive.
func Lookup(country, codeOrName string) (Subdivision, bool) {
	loadOnce.Do(load)
	idx := byCountry[strings.ToUpper(strings.TrimSpace(country))]
	key := fold(codeOrName)
	if idx == nil || key == "" {
		return Subdivision{}, false
	}
	if i, ok := idx.byCode[key]; ok {
		return idx.list[i], true
	}
	if i, ok := idx.byName[key]; ok && i >= 0 {
		return idx.list[i], true
	}
	return Subdivision{}, false
}

// fold reduces a code or name to lower-case letters and digits, without accents.
func fold(s string) string {
	var b strings.Builder
	for _, r := range norm.NFD.String(s) {
		switch {
		case unicode.Is(unicode.Mn, r):
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
