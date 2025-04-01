package country

var List map[string]Country

func init() {
	List = make(map[string]Country)

	List["AF"] = CommonCountry{
		name:         "Afghanistan",
		iso2:         "AF",
		iso3:         "AFG",
		numericCode:  4,
		dialPrefix:   93,
		currencyCode: "AFN",
	}
	List["AL"] = CommonCountry{
		name:         "Albania",
		iso2:         "AL",
		iso3:         "ALB",
		numericCode:  8,
		dialPrefix:   355,
		currencyCode: "ALL",
	}
	List["DZ"] = CommonCountry{
		name:         "Algeria",
		iso2:         "DZ",
		iso3:         "DZA",
		numericCode:  12,
		dialPrefix:   213,
		currencyCode: "DZD",
	}
	List["AS"] = CommonCountry{
		name:         "American Samoa",
		iso2:         "AS",
		iso3:         "ASM",
		numericCode:  16,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["AD"] = CommonCountry{
		name:         "Andorra",
		iso2:         "AD",
		iso3:         "AND",
		numericCode:  20,
		dialPrefix:   376,
		currencyCode: "EUR",
	}
	List["AO"] = CommonCountry{
		name:         "Angola",
		iso2:         "AO",
		iso3:         "AGO",
		numericCode:  24,
		dialPrefix:   244,
		currencyCode: "AOA",
	}
	List["AI"] = CommonCountry{
		name:         "Anguilla",
		iso2:         "AI",
		iso3:         "AIA",
		numericCode:  660,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["AQ"] = CommonCountry{
		name:         "Antarctica",
		iso2:         "AQ",
		iso3:         "ATA",
		numericCode:  10,
		dialPrefix:   672,
		currencyCode: "",
	}
	List["AG"] = CommonCountry{
		name:         "Antigua and Barbuda",
		iso2:         "AG",
		iso3:         "ATG",
		numericCode:  28,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["AR"] = CommonCountry{
		name:         "Argentina",
		iso2:         "AR",
		iso3:         "ARG",
		numericCode:  32,
		dialPrefix:   54,
		currencyCode: "ARS",
	}
	List["AM"] = CommonCountry{
		name:         "Armenia",
		iso2:         "AM",
		iso3:         "ARM",
		numericCode:  51,
		dialPrefix:   374,
		currencyCode: "AMD",
	}
	List["AW"] = CommonCountry{
		name:         "Aruba",
		iso2:         "AW",
		iso3:         "ABW",
		numericCode:  533,
		dialPrefix:   297,
		currencyCode: "AWG",
	}
	List["AU"] = CommonCountry{
		name:         "Australia",
		iso2:         "AU",
		iso3:         "AUS",
		numericCode:  36,
		dialPrefix:   61,
		currencyCode: "AUD",
	}
	List["AT"] = CommonCountry{
		name:         "Austria",
		iso2:         "AT",
		iso3:         "AUT",
		numericCode:  40,
		dialPrefix:   43,
		currencyCode: "EUR",
	}
	List["AZ"] = CommonCountry{
		name:         "Azerbaijan",
		iso2:         "AZ",
		iso3:         "AZE",
		numericCode:  31,
		dialPrefix:   994,
		currencyCode: "AZN",
	}
	List["BS"] = CommonCountry{
		name:         "Bahamas",
		iso2:         "BS",
		iso3:         "BHS",
		numericCode:  44,
		dialPrefix:   1,
		currencyCode: "BSD",
	}
	List["BH"] = CommonCountry{
		name:         "Bahrain",
		iso2:         "BH",
		iso3:         "BHR",
		numericCode:  48,
		dialPrefix:   973,
		currencyCode: "BHD",
	}
	List["BD"] = CommonCountry{
		name:         "Bangladesh",
		iso2:         "BD",
		iso3:         "BGD",
		numericCode:  50,
		dialPrefix:   880,
		currencyCode: "BDT",
	}
	List["BB"] = CommonCountry{
		name:         "Barbados",
		iso2:         "BB",
		iso3:         "BRB",
		numericCode:  52,
		dialPrefix:   1,
		currencyCode: "BBD",
	}
	List["BY"] = CommonCountry{
		name:         "Belarus",
		iso2:         "BY",
		iso3:         "BLR",
		numericCode:  112,
		dialPrefix:   375,
		currencyCode: "BYR",
	}
	List["BE"] = CommonCountry{
		name:         "Belgium",
		iso2:         "BE",
		iso3:         "BEL",
		numericCode:  56,
		dialPrefix:   32,
		currencyCode: "EUR",
	}
	List["BZ"] = CommonCountry{
		name:         "Belize",
		iso2:         "BZ",
		iso3:         "BLZ",
		numericCode:  84,
		dialPrefix:   501,
		currencyCode: "BZD",
	}
	List["BJ"] = CommonCountry{
		name:         "Benin",
		iso2:         "BJ",
		iso3:         "BEN",
		numericCode:  204,
		dialPrefix:   229,
		currencyCode: "XOF",
	}
	List["BM"] = CommonCountry{
		name:         "Bermuda",
		iso2:         "BM",
		iso3:         "BMU",
		numericCode:  60,
		dialPrefix:   1,
		currencyCode: "BMD",
	}
	List["BT"] = CommonCountry{
		name:         "Bhutan",
		iso2:         "BT",
		iso3:         "BTN",
		numericCode:  64,
		dialPrefix:   975,
		currencyCode: "INR",
	}
	List["BO"] = CommonCountry{
		name:         "Bolivia, Plurinational State of",
		iso2:         "BO",
		iso3:         "BOL",
		numericCode:  68,
		dialPrefix:   591,
		currencyCode: "BOB",
	}
	List["BQ"] = CommonCountry{
		name:         "Bonaire, Sint Eustatius and Saba",
		iso2:         "BQ",
		iso3:         "BES",
		numericCode:  535,
		dialPrefix:   599,
		currencyCode: "USD",
	}
	List["BA"] = CommonCountry{
		name:         "Bosnia and Herzegovina",
		iso2:         "BA",
		iso3:         "BIH",
		numericCode:  70,
		dialPrefix:   387,
		currencyCode: "BAM",
	}
	List["BW"] = CommonCountry{
		name:         "Botswana",
		iso2:         "BW",
		iso3:         "BWA",
		numericCode:  72,
		dialPrefix:   267,
		currencyCode: "BWP",
	}
	List["BV"] = CommonCountry{
		name:         "Bouvet Island",
		iso2:         "BV",
		iso3:         "BVT",
		numericCode:  74,
		dialPrefix:   47,
		currencyCode: "NOK",
	}
	List["BR"] = CommonCountry{
		name:         "Brazil",
		iso2:         "BR",
		iso3:         "BRA",
		numericCode:  76,
		dialPrefix:   55,
		currencyCode: "BRL",
	}
	List["IO"] = CommonCountry{
		name:         "British Indian Ocean Territory",
		iso2:         "IO",
		iso3:         "IOT",
		numericCode:  86,
		dialPrefix:   246,
		currencyCode: "USD",
	}
	List["BN"] = CommonCountry{
		name:         "Brunei Darussalam",
		iso2:         "BN",
		iso3:         "BRN",
		numericCode:  96,
		dialPrefix:   673,
		currencyCode: "BND",
	}
	List["BG"] = CommonCountry{
		name:         "Bulgaria",
		iso2:         "BG",
		iso3:         "BGR",
		numericCode:  100,
		dialPrefix:   359,
		currencyCode: "BGN",
	}
	List["BF"] = CommonCountry{
		name:         "Burkina Faso",
		iso2:         "BF",
		iso3:         "BFA",
		numericCode:  854,
		dialPrefix:   226,
		currencyCode: "XOF",
	}
	List["BI"] = CommonCountry{
		name:         "Burundi",
		iso2:         "BI",
		iso3:         "BDI",
		numericCode:  108,
		dialPrefix:   257,
		currencyCode: "BIF",
	}
	List["KH"] = CommonCountry{
		name:         "Cambodia",
		iso2:         "KH",
		iso3:         "KHM",
		numericCode:  116,
		dialPrefix:   855,
		currencyCode: "KHR",
	}
	List["CM"] = CommonCountry{
		name:         "Cameroon",
		iso2:         "CM",
		iso3:         "CMR",
		numericCode:  120,
		dialPrefix:   237,
		currencyCode: "XAF",
	}
	List["CA"] = CommonCountry{
		name:         "Canada",
		iso2:         "CA",
		iso3:         "CAN",
		numericCode:  124,
		dialPrefix:   1,
		currencyCode: "CAD",
	}
	List["CV"] = CommonCountry{
		name:         "Cape Verde",
		iso2:         "CV",
		iso3:         "CPV",
		numericCode:  132,
		dialPrefix:   238,
		currencyCode: "CVE",
	}
	List["KY"] = CommonCountry{
		name:         "Cayman Islands",
		iso2:         "KY",
		iso3:         "CYM",
		numericCode:  136,
		dialPrefix:   1,
		currencyCode: "KYD",
	}
	List["CF"] = CommonCountry{
		name:         "Central African Republic",
		iso2:         "CF",
		iso3:         "CAF",
		numericCode:  140,
		dialPrefix:   236,
		currencyCode: "XAF",
	}
	List["TD"] = CommonCountry{
		name:         "Chad",
		iso2:         "TD",
		iso3:         "TCD",
		numericCode:  148,
		dialPrefix:   235,
		currencyCode: "XAF",
	}
	List["CL"] = CommonCountry{
		name:         "Chile",
		iso2:         "CL",
		iso3:         "CHL",
		numericCode:  152,
		dialPrefix:   56,
		currencyCode: "CLP",
	}
	List["CN"] = CommonCountry{
		name:         "China",
		iso2:         "CN",
		iso3:         "CHN",
		numericCode:  156,
		dialPrefix:   86,
		currencyCode: "CNY",
	}
	List["CX"] = CommonCountry{
		name:         "Christmas Island",
		iso2:         "CX",
		iso3:         "CXR",
		numericCode:  162,
		dialPrefix:   61,
		currencyCode: "AUD",
	}
	List["CC"] = CommonCountry{
		name:         "Cocos (Keeling) Islands",
		iso2:         "CC",
		iso3:         "CCK",
		numericCode:  166,
		dialPrefix:   61,
		currencyCode: "AUD",
	}
	List["CO"] = CommonCountry{
		name:         "Colombia",
		iso2:         "CO",
		iso3:         "COL",
		numericCode:  170,
		dialPrefix:   57,
		currencyCode: "COP",
	}
	List["KM"] = CommonCountry{
		name:         "Comoros",
		iso2:         "KM",
		iso3:         "COM",
		numericCode:  174,
		dialPrefix:   269,
		currencyCode: "KMF",
	}
	List["CG"] = CommonCountry{
		name:         "Congo",
		iso2:         "CG",
		iso3:         "COG",
		numericCode:  178,
		dialPrefix:   242,
		currencyCode: "XAF",
	}
	List["CD"] = CommonCountry{
		name:         "Congo, the Democratic Republic of the",
		iso2:         "CD",
		iso3:         "COD",
		numericCode:  180,
		dialPrefix:   243,
		currencyCode: "",
	}
	List["CK"] = CommonCountry{
		name:         "Cook Islands",
		iso2:         "CK",
		iso3:         "COK",
		numericCode:  184,
		dialPrefix:   682,
		currencyCode: "NZD",
	}
	List["CR"] = CommonCountry{
		name:         "Costa Rica",
		iso2:         "CR",
		iso3:         "CRI",
		numericCode:  188,
		dialPrefix:   506,
		currencyCode: "CRC",
	}
	List["HR"] = CommonCountry{
		name:         "Croatia",
		iso2:         "HR",
		iso3:         "HRV",
		numericCode:  191,
		dialPrefix:   385,
		currencyCode: "HRK",
	}
	List["CU"] = CommonCountry{
		name:         "Cuba",
		iso2:         "CU",
		iso3:         "CUB",
		numericCode:  192,
		dialPrefix:   53,
		currencyCode: "CUP",
	}
	List["CW"] = CommonCountry{
		name:         "Curaçao",
		iso2:         "CW",
		iso3:         "CUW",
		numericCode:  531,
		dialPrefix:   599,
		currencyCode: "ANG",
	}
	List["CY"] = CommonCountry{
		name:         "Cyprus",
		iso2:         "CY",
		iso3:         "CYP",
		numericCode:  196,
		dialPrefix:   357,
		currencyCode: "EUR",
	}
	List["CZ"] = CommonCountry{
		name:         "Czech Republic",
		iso2:         "CZ",
		iso3:         "CZE",
		numericCode:  203,
		dialPrefix:   420,
		currencyCode: "CZK",
	}
	List["CI"] = CommonCountry{
		name:         "Côte d'Ivoire",
		iso2:         "CI",
		iso3:         "CIV",
		numericCode:  384,
		dialPrefix:   225,
		currencyCode: "XOF",
	}
	List["DK"] = CommonCountry{
		name:         "Denmark",
		iso2:         "DK",
		iso3:         "DNK",
		numericCode:  208,
		dialPrefix:   45,
		currencyCode: "DKK",
	}
	List["DJ"] = CommonCountry{
		name:         "Djibouti",
		iso2:         "DJ",
		iso3:         "DJI",
		numericCode:  262,
		dialPrefix:   253,
		currencyCode: "DJF",
	}
	List["DM"] = CommonCountry{
		name:         "Dominica",
		iso2:         "DM",
		iso3:         "DMA",
		numericCode:  212,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["DO"] = CommonCountry{
		name:         "Dominican Republic",
		iso2:         "DO",
		iso3:         "DOM",
		numericCode:  214,
		dialPrefix:   1,
		currencyCode: "DOP",
	}
	List["EC"] = CommonCountry{
		name:         "Ecuador",
		iso2:         "EC",
		iso3:         "ECU",
		numericCode:  218,
		dialPrefix:   593,
		currencyCode: "USD",
	}
	List["EG"] = CommonCountry{
		name:         "Egypt",
		iso2:         "EG",
		iso3:         "EGY",
		numericCode:  818,
		dialPrefix:   20,
		currencyCode: "EGP",
	}
	List["SV"] = CommonCountry{
		name:         "El Salvador",
		iso2:         "SV",
		iso3:         "SLV",
		numericCode:  222,
		dialPrefix:   503,
		currencyCode: "USD",
	}
	List["GQ"] = CommonCountry{
		name:         "Equatorial Guinea",
		iso2:         "GQ",
		iso3:         "GNQ",
		numericCode:  226,
		dialPrefix:   240,
		currencyCode: "XAF",
	}
	List["ER"] = CommonCountry{
		name:         "Eritrea",
		iso2:         "ER",
		iso3:         "ERI",
		numericCode:  232,
		dialPrefix:   291,
		currencyCode: "ERN",
	}
	List["EE"] = CommonCountry{
		name:         "Estonia",
		iso2:         "EE",
		iso3:         "EST",
		numericCode:  233,
		dialPrefix:   372,
		currencyCode: "EUR",
	}
	List["ET"] = CommonCountry{
		name:         "Ethiopia",
		iso2:         "ET",
		iso3:         "ETH",
		numericCode:  231,
		dialPrefix:   251,
		currencyCode: "ETB",
	}
	List["FK"] = CommonCountry{
		name:         "Falkland Islands (Malvinas)",
		iso2:         "FK",
		iso3:         "FLK",
		numericCode:  238,
		dialPrefix:   500,
		currencyCode: "FKP",
	}
	List["FO"] = CommonCountry{
		name:         "Faroe Islands",
		iso2:         "FO",
		iso3:         "FRO",
		numericCode:  234,
		dialPrefix:   298,
		currencyCode: "DKK",
	}
	List["FJ"] = CommonCountry{
		name:         "Fiji",
		iso2:         "FJ",
		iso3:         "FJI",
		numericCode:  242,
		dialPrefix:   679,
		currencyCode: "FJD",
	}
	List["FI"] = CommonCountry{
		name:         "Finland",
		iso2:         "FI",
		iso3:         "FIN",
		numericCode:  246,
		dialPrefix:   358,
		currencyCode: "EUR",
	}
	List["FR"] = CommonCountry{
		name:         "France",
		iso2:         "FR",
		iso3:         "FRA",
		numericCode:  250,
		dialPrefix:   33,
		currencyCode: "EUR",
	}
	List["GF"] = CommonCountry{
		name:         "French Guiana",
		iso2:         "GF",
		iso3:         "GUF",
		numericCode:  254,
		dialPrefix:   594,
		currencyCode: "EUR",
	}
	List["PF"] = CommonCountry{
		name:         "French Polynesia",
		iso2:         "PF",
		iso3:         "PYF",
		numericCode:  258,
		dialPrefix:   689,
		currencyCode: "XPF",
	}
	List["TF"] = CommonCountry{
		name:         "French Southern Territories",
		iso2:         "TF",
		iso3:         "ATF",
		numericCode:  260,
		dialPrefix:   262,
		currencyCode: "EUR",
	}
	List["GA"] = CommonCountry{
		name:         "Gabon",
		iso2:         "GA",
		iso3:         "GAB",
		numericCode:  266,
		dialPrefix:   241,
		currencyCode: "XAF",
	}
	List["GM"] = CommonCountry{
		name:         "Gambia",
		iso2:         "GM",
		iso3:         "GMB",
		numericCode:  270,
		dialPrefix:   220,
		currencyCode: "GMD",
	}
	List["GE"] = CommonCountry{
		name:         "Georgia",
		iso2:         "GE",
		iso3:         "GEO",
		numericCode:  268,
		dialPrefix:   995,
		currencyCode: "GEL",
	}
	List["DE"] = CommonCountry{
		name:         "Germany",
		iso2:         "DE",
		iso3:         "DEU",
		numericCode:  276,
		dialPrefix:   49,
		currencyCode: "EUR",
	}
	List["GH"] = CommonCountry{
		name:         "Ghana",
		iso2:         "GH",
		iso3:         "GHA",
		numericCode:  288,
		dialPrefix:   233,
		currencyCode: "GHS",
	}
	List["GI"] = CommonCountry{
		name:         "Gibraltar",
		iso2:         "GI",
		iso3:         "GIB",
		numericCode:  292,
		dialPrefix:   350,
		currencyCode: "GIP",
	}
	List["GR"] = CommonCountry{
		name:         "Greece",
		iso2:         "GR",
		iso3:         "GRC",
		numericCode:  300,
		dialPrefix:   30,
		currencyCode: "EUR",
	}
	List["GL"] = CommonCountry{
		name:         "Greenland",
		iso2:         "GL",
		iso3:         "GRL",
		numericCode:  304,
		dialPrefix:   299,
		currencyCode: "DKK",
	}
	List["GD"] = CommonCountry{
		name:         "Grenada",
		iso2:         "GD",
		iso3:         "GRD",
		numericCode:  308,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["GP"] = CommonCountry{
		name:         "Guadeloupe",
		iso2:         "GP",
		iso3:         "GLP",
		numericCode:  312,
		dialPrefix:   590,
		currencyCode: "EUR",
	}
	List["GU"] = CommonCountry{
		name:         "Guam",
		iso2:         "GU",
		iso3:         "GUM",
		numericCode:  316,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["GT"] = CommonCountry{
		name:         "Guatemala",
		iso2:         "GT",
		iso3:         "GTM",
		numericCode:  320,
		dialPrefix:   502,
		currencyCode: "GTQ",
	}
	List["GG"] = CommonCountry{
		name:         "Guernsey",
		iso2:         "GG",
		iso3:         "GGY",
		numericCode:  831,
		dialPrefix:   44,
		currencyCode: "GBP",
	}
	List["GN"] = CommonCountry{
		name:         "Guinea",
		iso2:         "GN",
		iso3:         "GIN",
		numericCode:  324,
		dialPrefix:   224,
		currencyCode: "GNF",
	}
	List["GW"] = CommonCountry{
		name:         "Guinea-Bissau",
		iso2:         "GW",
		iso3:         "GNB",
		numericCode:  624,
		dialPrefix:   245,
		currencyCode: "XOF",
	}
	List["GY"] = CommonCountry{
		name:         "Guyana",
		iso2:         "GY",
		iso3:         "GUY",
		numericCode:  328,
		dialPrefix:   592,
		currencyCode: "GYD",
	}
	List["HT"] = CommonCountry{
		name:         "Haiti",
		iso2:         "HT",
		iso3:         "HTI",
		numericCode:  332,
		dialPrefix:   509,
		currencyCode: "USD",
	}
	List["HM"] = CommonCountry{
		name:         "Heard Island and McDonald Islands",
		iso2:         "HM",
		iso3:         "HMD",
		numericCode:  334,
		dialPrefix:   672,
		currencyCode: "AUD",
	}
	List["VA"] = CommonCountry{
		name:         "Holy See (Vatican City State)",
		iso2:         "VA",
		iso3:         "VAT",
		numericCode:  336,
		dialPrefix:   39,
		currencyCode: "EUR",
	}
	List["HN"] = CommonCountry{
		name:         "Honduras",
		iso2:         "HN",
		iso3:         "HND",
		numericCode:  340,
		dialPrefix:   504,
		currencyCode: "HNL",
	}
	List["HK"] = CommonCountry{
		name:         "Hong Kong",
		iso2:         "HK",
		iso3:         "HKG",
		numericCode:  344,
		dialPrefix:   852,
		currencyCode: "HKD",
	}
	List["HU"] = CommonCountry{
		name:         "Hungary",
		iso2:         "HU",
		iso3:         "HUN",
		numericCode:  348,
		dialPrefix:   36,
		currencyCode: "HUF",
	}
	List["IS"] = CommonCountry{
		name:         "Iceland",
		iso2:         "IS",
		iso3:         "ISL",
		numericCode:  352,
		dialPrefix:   354,
		currencyCode: "ISK",
	}
	List["IN"] = CommonCountry{
		name:         "India",
		iso2:         "IN",
		iso3:         "IND",
		numericCode:  356,
		dialPrefix:   91,
		currencyCode: "INR",
	}
	List["ID"] = CommonCountry{
		name:         "Indonesia",
		iso2:         "ID",
		iso3:         "IDN",
		numericCode:  360,
		dialPrefix:   62,
		currencyCode: "IDR",
	}
	List["IR"] = CommonCountry{
		name:         "Iran, Islamic Republic of",
		iso2:         "IR",
		iso3:         "IRN",
		numericCode:  364,
		dialPrefix:   98,
		currencyCode: "IRR",
	}
	List["IQ"] = CommonCountry{
		name:         "Iraq",
		iso2:         "IQ",
		iso3:         "IRQ",
		numericCode:  368,
		dialPrefix:   964,
		currencyCode: "IQD",
	}
	List["IE"] = CommonCountry{
		name:         "Ireland",
		iso2:         "IE",
		iso3:         "IRL",
		numericCode:  372,
		dialPrefix:   353,
		currencyCode: "EUR",
	}
	List["IM"] = CommonCountry{
		name:         "Isle of Man",
		iso2:         "IM",
		iso3:         "IMN",
		numericCode:  833,
		dialPrefix:   44,
		currencyCode: "GBP",
	}
	List["IL"] = CommonCountry{
		name:         "Israel",
		iso2:         "IL",
		iso3:         "ISR",
		numericCode:  376,
		dialPrefix:   972,
		currencyCode: "ILS",
	}
	List["IT"] = CommonCountry{
		name:         "Italy",
		iso2:         "IT",
		iso3:         "ITA",
		numericCode:  380,
		dialPrefix:   39,
		currencyCode: "EUR",
	}
	List["JM"] = CommonCountry{
		name:         "Jamaica",
		iso2:         "JM",
		iso3:         "JAM",
		numericCode:  388,
		dialPrefix:   1,
		currencyCode: "JMD",
	}
	List["JP"] = CommonCountry{
		name:         "Japan",
		iso2:         "JP",
		iso3:         "JPN",
		numericCode:  392,
		dialPrefix:   81,
		currencyCode: "JPY",
	}
	List["JE"] = CommonCountry{
		name:         "Jersey",
		iso2:         "JE",
		iso3:         "JEY",
		numericCode:  832,
		dialPrefix:   44,
		currencyCode: "GBP",
	}
	List["JO"] = CommonCountry{
		name:         "Jordan",
		iso2:         "JO",
		iso3:         "JOR",
		numericCode:  400,
		dialPrefix:   962,
		currencyCode: "JOD",
	}
	List["KZ"] = CommonCountry{
		name:         "Kazakhstan",
		iso2:         "KZ",
		iso3:         "KAZ",
		numericCode:  398,
		dialPrefix:   7,
		currencyCode: "KZT",
	}
	List["KE"] = CommonCountry{
		name:         "Kenya",
		iso2:         "KE",
		iso3:         "KEN",
		numericCode:  404,
		dialPrefix:   254,
		currencyCode: "KES",
	}
	List["KI"] = CommonCountry{
		name:         "Kiribati",
		iso2:         "KI",
		iso3:         "KIR",
		numericCode:  296,
		dialPrefix:   686,
		currencyCode: "AUD",
	}
	List["KP"] = CommonCountry{
		name:         "Korea, Democratic People's Republic of",
		iso2:         "KP",
		iso3:         "PRK",
		numericCode:  408,
		dialPrefix:   850,
		currencyCode: "KPW",
	}
	List["KR"] = CommonCountry{
		name:         "Korea, Republic of",
		iso2:         "KR",
		iso3:         "KOR",
		numericCode:  410,
		dialPrefix:   82,
		currencyCode: "KRW",
	}
	List["KW"] = CommonCountry{
		name:         "Kuwait",
		iso2:         "KW",
		iso3:         "KWT",
		numericCode:  414,
		dialPrefix:   965,
		currencyCode: "KWD",
	}
	List["KG"] = CommonCountry{
		name:         "Kyrgyzstan",
		iso2:         "KG",
		iso3:         "KGZ",
		numericCode:  417,
		dialPrefix:   996,
		currencyCode: "KGS",
	}
	List["LA"] = CommonCountry{
		name:         "Lao People's Democratic Republic",
		iso2:         "LA",
		iso3:         "LAO",
		numericCode:  418,
		dialPrefix:   856,
		currencyCode: "LAK",
	}
	List["LV"] = CommonCountry{
		name:         "Latvia",
		iso2:         "LV",
		iso3:         "LVA",
		numericCode:  428,
		dialPrefix:   371,
		currencyCode: "EUR",
	}
	List["LB"] = CommonCountry{
		name:         "Lebanon",
		iso2:         "LB",
		iso3:         "LBN",
		numericCode:  422,
		dialPrefix:   961,
		currencyCode: "LBP",
	}
	List["LS"] = CommonCountry{
		name:         "Lesotho",
		iso2:         "LS",
		iso3:         "LSO",
		numericCode:  426,
		dialPrefix:   266,
		currencyCode: "ZAR",
	}
	List["LR"] = CommonCountry{
		name:         "Liberia",
		iso2:         "LR",
		iso3:         "LBR",
		numericCode:  430,
		dialPrefix:   231,
		currencyCode: "LRD",
	}
	List["LY"] = CommonCountry{
		name:         "Libya",
		iso2:         "LY",
		iso3:         "LBY",
		numericCode:  434,
		dialPrefix:   218,
		currencyCode: "LYD",
	}
	List["LI"] = CommonCountry{
		name:         "Liechtenstein",
		iso2:         "LI",
		iso3:         "LIE",
		numericCode:  438,
		dialPrefix:   423,
		currencyCode: "CHF",
	}
	List["LT"] = CommonCountry{
		name:         "Lithuania",
		iso2:         "LT",
		iso3:         "LTU",
		numericCode:  440,
		dialPrefix:   370,
		currencyCode: "EUR",
	}
	List["LU"] = CommonCountry{
		name:         "Luxembourg",
		iso2:         "LU",
		iso3:         "LUX",
		numericCode:  442,
		dialPrefix:   352,
		currencyCode: "EUR",
	}
	List["MO"] = CommonCountry{
		name:         "Macao",
		iso2:         "MO",
		iso3:         "MAC",
		numericCode:  446,
		dialPrefix:   853,
		currencyCode: "MOP",
	}
	List["MK"] = CommonCountry{
		name:         "Macedonia, the Former Yugoslav Republic of",
		iso2:         "MK",
		iso3:         "MKD",
		numericCode:  807,
		dialPrefix:   389,
		currencyCode: "MKD",
	}
	List["MG"] = CommonCountry{
		name:         "Madagascar",
		iso2:         "MG",
		iso3:         "MDG",
		numericCode:  450,
		dialPrefix:   261,
		currencyCode: "MGA",
	}
	List["MW"] = CommonCountry{
		name:         "Malawi",
		iso2:         "MW",
		iso3:         "MWI",
		numericCode:  454,
		dialPrefix:   265,
		currencyCode: "MWK",
	}
	List["MY"] = CommonCountry{
		name:         "Malaysia",
		iso2:         "MY",
		iso3:         "MYS",
		numericCode:  458,
		dialPrefix:   60,
		currencyCode: "MYR",
	}
	List["MV"] = CommonCountry{
		name:         "Maldives",
		iso2:         "MV",
		iso3:         "MDV",
		numericCode:  462,
		dialPrefix:   960,
		currencyCode: "MVR",
	}
	List["ML"] = CommonCountry{
		name:         "Mali",
		iso2:         "ML",
		iso3:         "MLI",
		numericCode:  466,
		dialPrefix:   223,
		currencyCode: "XOF",
	}
	List["MT"] = CommonCountry{
		name:         "Malta",
		iso2:         "MT",
		iso3:         "MLT",
		numericCode:  470,
		dialPrefix:   356,
		currencyCode: "EUR",
	}
	List["MH"] = CommonCountry{
		name:         "Marshall Islands",
		iso2:         "MH",
		iso3:         "MHL",
		numericCode:  584,
		dialPrefix:   692,
		currencyCode: "USD",
	}
	List["MQ"] = CommonCountry{
		name:         "Martinique",
		iso2:         "MQ",
		iso3:         "MTQ",
		numericCode:  474,
		dialPrefix:   596,
		currencyCode: "EUR",
	}
	List["MR"] = CommonCountry{
		name:         "Mauritania",
		iso2:         "MR",
		iso3:         "MRT",
		numericCode:  478,
		dialPrefix:   222,
		currencyCode: "MRO",
	}
	List["MU"] = CommonCountry{
		name:         "Mauritius",
		iso2:         "MU",
		iso3:         "MUS",
		numericCode:  480,
		dialPrefix:   230,
		currencyCode: "MUR",
	}
	List["YT"] = CommonCountry{
		name:         "Mayotte",
		iso2:         "YT",
		iso3:         "MYT",
		numericCode:  175,
		dialPrefix:   262,
		currencyCode: "EUR",
	}
	List["MX"] = CommonCountry{
		name:         "Mexico",
		iso2:         "MX",
		iso3:         "MEX",
		numericCode:  484,
		dialPrefix:   52,
		currencyCode: "MXN",
	}
	List["FM"] = CommonCountry{
		name:         "Micronesia, Federated States of",
		iso2:         "FM",
		iso3:         "FSM",
		numericCode:  583,
		dialPrefix:   691,
		currencyCode: "USD",
	}
	List["MD"] = CommonCountry{
		name:         "Moldova, Republic of",
		iso2:         "MD",
		iso3:         "MDA",
		numericCode:  498,
		dialPrefix:   373,
		currencyCode: "MDL",
	}
	List["MC"] = CommonCountry{
		name:         "Monaco",
		iso2:         "MC",
		iso3:         "MCO",
		numericCode:  492,
		dialPrefix:   377,
		currencyCode: "EUR",
	}
	List["MN"] = CommonCountry{
		name:         "Mongolia",
		iso2:         "MN",
		iso3:         "MNG",
		numericCode:  496,
		dialPrefix:   976,
		currencyCode: "MNT",
	}
	List["ME"] = CommonCountry{
		name:         "Montenegro",
		iso2:         "ME",
		iso3:         "MNE",
		numericCode:  499,
		dialPrefix:   382,
		currencyCode: "EUR",
	}
	List["MS"] = CommonCountry{
		name:         "Montserrat",
		iso2:         "MS",
		iso3:         "MSR",
		numericCode:  500,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["MA"] = CommonCountry{
		name:         "Morocco",
		iso2:         "MA",
		iso3:         "MAR",
		numericCode:  504,
		dialPrefix:   212,
		currencyCode: "MAD",
	}
	List["MZ"] = CommonCountry{
		name:         "Mozambique",
		iso2:         "MZ",
		iso3:         "MOZ",
		numericCode:  508,
		dialPrefix:   258,
		currencyCode: "MZN",
	}
	List["MM"] = CommonCountry{
		name:         "Myanmar",
		iso2:         "MM",
		iso3:         "MMR",
		numericCode:  104,
		dialPrefix:   95,
		currencyCode: "MMK",
	}
	List["NA"] = CommonCountry{
		name:         "Namibia",
		iso2:         "NA",
		iso3:         "NAM",
		numericCode:  516,
		dialPrefix:   264,
		currencyCode: "ZAR",
	}
	List["NR"] = CommonCountry{
		name:         "Nauru",
		iso2:         "NR",
		iso3:         "NRU",
		numericCode:  520,
		dialPrefix:   674,
		currencyCode: "AUD",
	}
	List["NP"] = CommonCountry{
		name:         "Nepal",
		iso2:         "NP",
		iso3:         "NPL",
		numericCode:  524,
		dialPrefix:   977,
		currencyCode: "NPR",
	}
	List["NL"] = CommonCountry{
		name:         "Netherlands",
		iso2:         "NL",
		iso3:         "NLD",
		numericCode:  528,
		dialPrefix:   31,
		currencyCode: "EUR",
	}
	List["NC"] = CommonCountry{
		name:         "New Caledonia",
		iso2:         "NC",
		iso3:         "NCL",
		numericCode:  540,
		dialPrefix:   687,
		currencyCode: "XPF",
	}
	List["NZ"] = CommonCountry{
		name:         "New Zealand",
		iso2:         "NZ",
		iso3:         "NZL",
		numericCode:  554,
		dialPrefix:   64,
		currencyCode: "NZD",
	}
	List["NI"] = CommonCountry{
		name:         "Nicaragua",
		iso2:         "NI",
		iso3:         "NIC",
		numericCode:  558,
		dialPrefix:   505,
		currencyCode: "NIO",
	}
	List["NE"] = CommonCountry{
		name:         "Niger",
		iso2:         "NE",
		iso3:         "NER",
		numericCode:  562,
		dialPrefix:   227,
		currencyCode: "XOF",
	}
	List["NG"] = CommonCountry{
		name:         "Nigeria",
		iso2:         "NG",
		iso3:         "NGA",
		numericCode:  566,
		dialPrefix:   234,
		currencyCode: "NGN",
	}
	List["NU"] = CommonCountry{
		name:         "Niue",
		iso2:         "NU",
		iso3:         "NIU",
		numericCode:  570,
		dialPrefix:   683,
		currencyCode: "NZD",
	}
	List["NF"] = CommonCountry{
		name:         "Norfolk Island",
		iso2:         "NF",
		iso3:         "NFK",
		numericCode:  574,
		dialPrefix:   672,
		currencyCode: "AUD",
	}
	List["MP"] = CommonCountry{
		name:         "Northern Mariana Islands",
		iso2:         "MP",
		iso3:         "MNP",
		numericCode:  580,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["NO"] = CommonCountry{
		name:         "Norway",
		iso2:         "NO",
		iso3:         "NOR",
		numericCode:  578,
		dialPrefix:   47,
		currencyCode: "NOK",
	}
	List["OM"] = CommonCountry{
		name:         "Oman",
		iso2:         "OM",
		iso3:         "OMN",
		numericCode:  512,
		dialPrefix:   968,
		currencyCode: "OMR",
	}
	List["PK"] = CommonCountry{
		name:         "Pakistan",
		iso2:         "PK",
		iso3:         "PAK",
		numericCode:  586,
		dialPrefix:   92,
		currencyCode: "PKR",
	}
	List["PW"] = CommonCountry{
		name:         "Palau",
		iso2:         "PW",
		iso3:         "PLW",
		numericCode:  585,
		dialPrefix:   680,
		currencyCode: "USD",
	}
	List["PS"] = CommonCountry{
		name:         "Palestine, State of",
		iso2:         "PS",
		iso3:         "PSE",
		numericCode:  275,
		dialPrefix:   970,
		currencyCode: "",
	}
	List["PA"] = CommonCountry{
		name:         "Panama",
		iso2:         "PA",
		iso3:         "PAN",
		numericCode:  591,
		dialPrefix:   507,
		currencyCode: "USD",
	}
	List["PG"] = CommonCountry{
		name:         "Papua New Guinea",
		iso2:         "PG",
		iso3:         "PNG",
		numericCode:  598,
		dialPrefix:   675,
		currencyCode: "PGK",
	}
	List["PY"] = CommonCountry{
		name:         "Paraguay",
		iso2:         "PY",
		iso3:         "PRY",
		numericCode:  600,
		dialPrefix:   595,
		currencyCode: "PYG",
	}
	List["PE"] = CommonCountry{
		name:         "Peru",
		iso2:         "PE",
		iso3:         "PER",
		numericCode:  604,
		dialPrefix:   51,
		currencyCode: "PEN",
	}
	List["PH"] = CommonCountry{
		name:         "Philippines",
		iso2:         "PH",
		iso3:         "PHL",
		numericCode:  608,
		dialPrefix:   63,
		currencyCode: "PHP",
	}
	List["PN"] = CommonCountry{
		name:         "Pitcairn",
		iso2:         "PN",
		iso3:         "PCN",
		numericCode:  612,
		dialPrefix:   870,
		currencyCode: "NZD",
	}
	List["PL"] = CommonCountry{
		name:         "Poland",
		iso2:         "PL",
		iso3:         "POL",
		numericCode:  616,
		dialPrefix:   48,
		currencyCode: "PLN",
	}
	List["PT"] = CommonCountry{
		name:         "Portugal",
		iso2:         "PT",
		iso3:         "PRT",
		numericCode:  620,
		dialPrefix:   351,
		currencyCode: "EUR",
	}
	List["PR"] = CommonCountry{
		name:         "Puerto Rico",
		iso2:         "PR",
		iso3:         "PRI",
		numericCode:  630,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["QA"] = CommonCountry{
		name:         "Qatar",
		iso2:         "QA",
		iso3:         "QAT",
		numericCode:  634,
		dialPrefix:   974,
		currencyCode: "QAR",
	}
	List["RO"] = CommonCountry{
		name:         "Romania",
		iso2:         "RO",
		iso3:         "ROU",
		numericCode:  642,
		dialPrefix:   40,
		currencyCode: "RON",
	}
	List["RU"] = CommonCountry{
		name:         "Russian Federation",
		iso2:         "RU",
		iso3:         "RUS",
		numericCode:  643,
		dialPrefix:   7,
		currencyCode: "RUB",
	}
	List["RW"] = CommonCountry{
		name:         "Rwanda",
		iso2:         "RW",
		iso3:         "RWA",
		numericCode:  646,
		dialPrefix:   250,
		currencyCode: "RWF",
	}
	List["RE"] = CommonCountry{
		name:         "Réunion",
		iso2:         "RE",
		iso3:         "REU",
		numericCode:  638,
		dialPrefix:   262,
		currencyCode: "EUR",
	}
	List["BL"] = CommonCountry{
		name:         "Saint Barthélemy",
		iso2:         "BL",
		iso3:         "BLM",
		numericCode:  652,
		dialPrefix:   590,
		currencyCode: "EUR",
	}
	List["SH"] = CommonCountry{
		name:         "Saint Helena, Ascension and Tristan da Cunha",
		iso2:         "SH",
		iso3:         "SHN",
		numericCode:  654,
		dialPrefix:   290,
		currencyCode: "SHP",
	}
	List["KN"] = CommonCountry{
		name:         "Saint Kitts and Nevis",
		iso2:         "KN",
		iso3:         "KNA",
		numericCode:  659,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["LC"] = CommonCountry{
		name:         "Saint Lucia",
		iso2:         "LC",
		iso3:         "LCA",
		numericCode:  662,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["MF"] = CommonCountry{
		name:         "Saint Martin (French part)",
		iso2:         "MF",
		iso3:         "MAF",
		numericCode:  663,
		dialPrefix:   590,
		currencyCode: "EUR",
	}
	List["PM"] = CommonCountry{
		name:         "Saint Pierre and Miquelon",
		iso2:         "PM",
		iso3:         "SPM",
		numericCode:  666,
		dialPrefix:   508,
		currencyCode: "EUR",
	}
	List["VC"] = CommonCountry{
		name:         "Saint Vincent and the Grenadines",
		iso2:         "VC",
		iso3:         "VCT",
		numericCode:  670,
		dialPrefix:   1,
		currencyCode: "XCD",
	}
	List["WS"] = CommonCountry{
		name:         "Samoa",
		iso2:         "WS",
		iso3:         "WSM",
		numericCode:  882,
		dialPrefix:   685,
		currencyCode: "WST",
	}
	List["SM"] = CommonCountry{
		name:         "San Marino",
		iso2:         "SM",
		iso3:         "SMR",
		numericCode:  674,
		dialPrefix:   378,
		currencyCode: "EUR",
	}
	List["ST"] = CommonCountry{
		name:         "Sao Tome and Principe",
		iso2:         "ST",
		iso3:         "STP",
		numericCode:  678,
		dialPrefix:   239,
		currencyCode: "STD",
	}
	List["SA"] = CommonCountry{
		name:         "Saudi Arabia",
		iso2:         "SA",
		iso3:         "SAU",
		numericCode:  682,
		dialPrefix:   966,
		currencyCode: "SAR",
	}
	List["SN"] = CommonCountry{
		name:         "Senegal",
		iso2:         "SN",
		iso3:         "SEN",
		numericCode:  686,
		dialPrefix:   221,
		currencyCode: "XOF",
	}
	List["RS"] = CommonCountry{
		name:         "Serbia",
		iso2:         "RS",
		iso3:         "SRB",
		numericCode:  688,
		dialPrefix:   381,
		currencyCode: "RSD",
	}
	List["SC"] = CommonCountry{
		name:         "Seychelles",
		iso2:         "SC",
		iso3:         "SYC",
		numericCode:  690,
		dialPrefix:   248,
		currencyCode: "SCR",
	}
	List["SL"] = CommonCountry{
		name:         "Sierra Leone",
		iso2:         "SL",
		iso3:         "SLE",
		numericCode:  694,
		dialPrefix:   232,
		currencyCode: "SLL",
	}
	List["SG"] = CommonCountry{
		name:         "Singapore",
		iso2:         "SG",
		iso3:         "SGP",
		numericCode:  702,
		dialPrefix:   65,
		currencyCode: "SGD",
	}
	List["SX"] = CommonCountry{
		name:         "Sint Maarten (Dutch part)",
		iso2:         "SX",
		iso3:         "SXM",
		numericCode:  534,
		dialPrefix:   1,
		currencyCode: "ANG",
	}
	List["SK"] = CommonCountry{
		name:         "Slovakia",
		iso2:         "SK",
		iso3:         "SVK",
		numericCode:  703,
		dialPrefix:   421,
		currencyCode: "EUR",
	}
	List["SI"] = CommonCountry{
		name:         "Slovenia",
		iso2:         "SI",
		iso3:         "SVN",
		numericCode:  705,
		dialPrefix:   386,
		currencyCode: "EUR",
	}
	List["SB"] = CommonCountry{
		name:         "Solomon Islands",
		iso2:         "SB",
		iso3:         "SLB",
		numericCode:  90,
		dialPrefix:   677,
		currencyCode: "SBD",
	}
	List["SO"] = CommonCountry{
		name:         "Somalia",
		iso2:         "SO",
		iso3:         "SOM",
		numericCode:  706,
		dialPrefix:   252,
		currencyCode: "SOS",
	}
	List["ZA"] = CommonCountry{
		name:         "South Africa",
		iso2:         "ZA",
		iso3:         "ZAF",
		numericCode:  710,
		dialPrefix:   27,
		currencyCode: "ZAR",
	}
	List["GS"] = CommonCountry{
		name:         "South Georgia and the South Sandwich Islands",
		iso2:         "GS",
		iso3:         "SGS",
		numericCode:  239,
		dialPrefix:   500,
		currencyCode: "",
	}
	List["SS"] = CommonCountry{
		name:         "South Sudan",
		iso2:         "SS",
		iso3:         "SSD",
		numericCode:  728,
		dialPrefix:   211,
		currencyCode: "SSP",
	}
	List["ES"] = CommonCountry{
		name:         "Spain",
		iso2:         "ES",
		iso3:         "ESP",
		numericCode:  724,
		dialPrefix:   34,
		currencyCode: "EUR",
	}
	List["LK"] = CommonCountry{
		name:         "Sri Lanka",
		iso2:         "LK",
		iso3:         "LKA",
		numericCode:  144,
		dialPrefix:   94,
		currencyCode: "LKR",
	}
	List["SD"] = CommonCountry{
		name:         "Sudan",
		iso2:         "SD",
		iso3:         "SDN",
		numericCode:  729,
		dialPrefix:   249,
		currencyCode: "SDG",
	}
	List["SR"] = CommonCountry{
		name:         "Suriname",
		iso2:         "SR",
		iso3:         "SUR",
		numericCode:  740,
		dialPrefix:   597,
		currencyCode: "SRD",
	}
	List["SJ"] = CommonCountry{
		name:         "Svalbard and Jan Mayen",
		iso2:         "SJ",
		iso3:         "SJM",
		numericCode:  744,
		dialPrefix:   47,
		currencyCode: "NOK",
	}
	List["SZ"] = CommonCountry{
		name:         "Swaziland",
		iso2:         "SZ",
		iso3:         "SWZ",
		numericCode:  748,
		dialPrefix:   268,
		currencyCode: "SZL",
	}
	List["SE"] = CommonCountry{
		name:         "Sweden",
		iso2:         "SE",
		iso3:         "SWE",
		numericCode:  752,
		dialPrefix:   46,
		currencyCode: "SEK",
	}
	List["CH"] = CommonCountry{
		name:         "Switzerland",
		iso2:         "CH",
		iso3:         "CHE",
		numericCode:  756,
		dialPrefix:   41,
		currencyCode: "CHF",
	}
	List["SY"] = CommonCountry{
		name:         "Syrian Arab Republic",
		iso2:         "SY",
		iso3:         "SYR",
		numericCode:  760,
		dialPrefix:   963,
		currencyCode: "SYP",
	}
	List["TW"] = CommonCountry{
		name:         "Taiwan, Province of China",
		iso2:         "TW",
		iso3:         "TWN",
		numericCode:  158,
		dialPrefix:   886,
		currencyCode: "TWD",
	}
	List["TJ"] = CommonCountry{
		name:         "Tajikistan",
		iso2:         "TJ",
		iso3:         "TJK",
		numericCode:  762,
		dialPrefix:   992,
		currencyCode: "TJS",
	}
	List["TZ"] = CommonCountry{
		name:         "Tanzania, United Republic of",
		iso2:         "TZ",
		iso3:         "TZA",
		numericCode:  834,
		dialPrefix:   255,
		currencyCode: "TZS",
	}
	List["TH"] = CommonCountry{
		name:         "Thailand",
		iso2:         "TH",
		iso3:         "THA",
		numericCode:  764,
		dialPrefix:   66,
		currencyCode: "THB",
	}
	List["TL"] = CommonCountry{
		name:         "Timor-Leste",
		iso2:         "TL",
		iso3:         "TLS",
		numericCode:  626,
		dialPrefix:   670,
		currencyCode: "USD",
	}
	List["TG"] = CommonCountry{
		name:         "Togo",
		iso2:         "TG",
		iso3:         "TGO",
		numericCode:  768,
		dialPrefix:   228,
		currencyCode: "XOF",
	}
	List["TK"] = CommonCountry{
		name:         "Tokelau",
		iso2:         "TK",
		iso3:         "TKL",
		numericCode:  772,
		dialPrefix:   690,
		currencyCode: "NZD",
	}
	List["TO"] = CommonCountry{
		name:         "Tonga",
		iso2:         "TO",
		iso3:         "TON",
		numericCode:  776,
		dialPrefix:   676,
		currencyCode: "TOP",
	}
	List["TT"] = CommonCountry{
		name:         "Trinidad and Tobago",
		iso2:         "TT",
		iso3:         "TTO",
		numericCode:  780,
		dialPrefix:   1,
		currencyCode: "TTD",
	}
	List["TN"] = CommonCountry{
		name:         "Tunisia",
		iso2:         "TN",
		iso3:         "TUN",
		numericCode:  788,
		dialPrefix:   216,
		currencyCode: "TND",
	}
	List["TR"] = CommonCountry{
		name:         "Turkey",
		iso2:         "TR",
		iso3:         "TUR",
		numericCode:  792,
		dialPrefix:   90,
		currencyCode: "TRY",
	}
	List["TM"] = CommonCountry{
		name:         "Turkmenistan",
		iso2:         "TM",
		iso3:         "TKM",
		numericCode:  795,
		dialPrefix:   993,
		currencyCode: "TMT",
	}
	List["TC"] = CommonCountry{
		name:         "Turks and Caicos Islands",
		iso2:         "TC",
		iso3:         "TCA",
		numericCode:  796,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["TV"] = CommonCountry{
		name:         "Tuvalu",
		iso2:         "TV",
		iso3:         "TUV",
		numericCode:  798,
		dialPrefix:   688,
		currencyCode: "AUD",
	}
	List["UG"] = CommonCountry{
		name:         "Uganda",
		iso2:         "UG",
		iso3:         "UGA",
		numericCode:  800,
		dialPrefix:   256,
		currencyCode: "UGX",
	}
	List["UA"] = CommonCountry{
		name:         "Ukraine",
		iso2:         "UA",
		iso3:         "UKR",
		numericCode:  804,
		dialPrefix:   380,
		currencyCode: "UAH",
	}
	List["AE"] = CommonCountry{
		name:         "United Arab Emirates",
		iso2:         "AE",
		iso3:         "ARE",
		numericCode:  784,
		dialPrefix:   971,
		currencyCode: "AED",
	}
	List["GB"] = CommonCountry{
		name:         "United Kingdom",
		iso2:         "GB",
		iso3:         "GBR",
		numericCode:  826,
		dialPrefix:   44,
		currencyCode: "GBP",
	}
	List["US"] = CommonCountry{
		name:         "United States",
		iso2:         "US",
		iso3:         "USA",
		numericCode:  840,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["UM"] = CommonCountry{
		name:         "United States Minor Outlying Islands",
		iso2:         "UM",
		iso3:         "UMI",
		numericCode:  581,
		dialPrefix:   0,
		currencyCode: "USD",
	}
	List["UY"] = CommonCountry{
		name:         "Uruguay",
		iso2:         "UY",
		iso3:         "URY",
		numericCode:  858,
		dialPrefix:   598,
		currencyCode: "UYU",
	}
	List["UZ"] = CommonCountry{
		name:         "Uzbekistan",
		iso2:         "UZ",
		iso3:         "UZB",
		numericCode:  860,
		dialPrefix:   998,
		currencyCode: "UZS",
	}
	List["VU"] = CommonCountry{
		name:         "Vanuatu",
		iso2:         "VU",
		iso3:         "VUT",
		numericCode:  548,
		dialPrefix:   678,
		currencyCode: "VUV",
	}
	List["VE"] = CommonCountry{
		name:         "Venezuela, Bolivarian Republic of",
		iso2:         "VE",
		iso3:         "VEN",
		numericCode:  862,
		dialPrefix:   58,
		currencyCode: "VEF",
	}
	List["VN"] = CommonCountry{
		name:         "Viet Nam",
		iso2:         "VN",
		iso3:         "VNM",
		numericCode:  704,
		dialPrefix:   84,
		currencyCode: "VND",
	}
	List["VG"] = CommonCountry{
		name:         "Virgin Islands, British",
		iso2:         "VG",
		iso3:         "VGB",
		numericCode:  92,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["VI"] = CommonCountry{
		name:         "Virgin Islands, U.S.",
		iso2:         "VI",
		iso3:         "VIR",
		numericCode:  850,
		dialPrefix:   1,
		currencyCode: "USD",
	}
	List["WF"] = CommonCountry{
		name:         "Wallis and Futuna",
		iso2:         "WF",
		iso3:         "WLF",
		numericCode:  876,
		dialPrefix:   681,
		currencyCode: "XPF",
	}
	List["EH"] = CommonCountry{
		name:         "Western Sahara",
		iso2:         "EH",
		iso3:         "ESH",
		numericCode:  732,
		dialPrefix:   212,
		currencyCode: "MAD",
	}
	List["YE"] = CommonCountry{
		name:         "Yemen",
		iso2:         "YE",
		iso3:         "YEM",
		numericCode:  887,
		dialPrefix:   967,
		currencyCode: "YER",
	}
	List["ZM"] = CommonCountry{
		name:         "Zambia",
		iso2:         "ZM",
		iso3:         "ZMB",
		numericCode:  894,
		dialPrefix:   260,
		currencyCode: "ZMW",
	}
	List["ZW"] = CommonCountry{
		name:         "Zimbabwe",
		iso2:         "ZW",
		iso3:         "ZWE",
		numericCode:  716,
		dialPrefix:   263,
		currencyCode: "ZWL",
	}
	List["AX"] = CommonCountry{
		name:         "Åland Islands",
		iso2:         "AX",
		iso3:         "ALA",
		numericCode:  248,
		dialPrefix:   358,
		currencyCode: "EUR",
	}
}
