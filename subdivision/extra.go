package subdivision

// extra are ISO 3166-2 subdivisions CLDR leaves out: the US outlying areas, which CLDR
// treats only as countries, though US addresses give them as states.
var extra = []Subdivision{
	{Code: "US-AS", Name: "American Samoa"},
	{Code: "US-GU", Name: "Guam"},
	{Code: "US-MP", Name: "Northern Mariana Islands"},
	{Code: "US-PR", Name: "Puerto Rico"},
	{Code: "US-UM", Name: "United States Minor Outlying Islands"},
	{Code: "US-VI", Name: "Virgin Islands, U.S."},
}

// aliases are other names subdivisions are given in addresses.
var aliases = map[string][]string{
	"US-DC": {"District of Columbia", "Washington D.C."},
	"US-VI": {"US Virgin Islands", "U.S. Virgin Islands", "United States Virgin Islands"},
	"CA-BC": {"Colombie-Britannique"},
	"CA-NB": {"Nouveau-Brunswick"},
	"CA-NL": {"Newfoundland", "Terre-Neuve-et-Labrador"},
	"CA-NS": {"Nouvelle-Écosse"},
	"CA-PE": {"Île-du-Prince-Édouard"},
	"CA-QC": {"Québec"},
	"CA-NT": {"Territoires du Nord-Ouest"},
	"CA-YT": {"Yukon Territory"},
}
