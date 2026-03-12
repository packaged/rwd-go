package currency

var Currencies map[string]Currency

func init() {
	Currencies = make(map[string]Currency)

	Currencies["AED"] = CommonCurrency{
		name:              "UAE Dirham",
		symbol:            "د.إ",
		code:              "AED",
		numericCode:       784,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dirham",
		minorUnit:         "fils",
		formatPattern:     "{amount} {code}",
	}

	Currencies["AFN"] = CommonCurrency{
		name:              "Afghani",
		symbol:            "؋",
		code:              "AFN",
		numericCode:       971,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "afghani",
		minorUnit:         "pul",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["ALL"] = CommonCurrency{
		name:              "Lek",
		symbol:            "L",
		code:              "ALL",
		numericCode:       8,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "lek",
		minorUnit:         "qindarkë",
		formatPattern:     "{amount} {code}",
	}

	Currencies["AMD"] = CommonCurrency{
		name:              "Armenian Dram",
		symbol:            "֏",
		code:              "AMD",
		numericCode:       51,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dram",
		minorUnit:         "luma",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["AOA"] = CommonCurrency{
		name:              "Kwanza",
		symbol:            "Kz",
		code:              "AOA",
		numericCode:       973,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kwanza",
		minorUnit:         "cêntimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["ARS"] = CommonCurrency{
		name:              "Argentine Peso",
		symbol:            "$",
		code:              "ARS",
		numericCode:       32,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["AUD"] = CommonCurrency{
		name:              "Australian Dollar",
		symbol:            "$",
		code:              "AUD",
		numericCode:       36,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["AWG"] = CommonCurrency{
		name:              "Aruban Florin",
		symbol:            "ƒ",
		code:              "AWG",
		numericCode:       533,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "florin",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["AZN"] = CommonCurrency{
		name:              "Azerbaijan Manat",
		symbol:            "₼",
		code:              "AZN",
		numericCode:       944,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "manat",
		minorUnit:         "qəpik",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BAM"] = CommonCurrency{
		name:              "Convertible Mark",
		symbol:            "KM",
		code:              "BAM",
		numericCode:       977,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "mark",
		minorUnit:         "fening",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["BBD"] = CommonCurrency{
		name:              "Barbados Dollar",
		symbol:            "$",
		code:              "BBD",
		numericCode:       52,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BDT"] = CommonCurrency{
		name:              "Taka",
		symbol:            "৳",
		code:              "BDT",
		numericCode:       50,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "taka",
		minorUnit:         "paisa",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BHD"] = CommonCurrency{
		name:              "Bahraini Dinar",
		symbol:            "BD",
		code:              "BHD",
		numericCode:       48,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "fils",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BIF"] = CommonCurrency{
		name:              "Burundi Franc",
		symbol:            "FBu",
		code:              "BIF",
		numericCode:       108,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["BMD"] = CommonCurrency{
		name:              "Bermudian Dollar",
		symbol:            "$",
		code:              "BMD",
		numericCode:       60,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BND"] = CommonCurrency{
		name:              "Brunei Dollar",
		symbol:            "$",
		code:              "BND",
		numericCode:       96,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "sen",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BOB"] = CommonCurrency{
		name:              "Boliviano",
		symbol:            "Bs.",
		code:              "BOB",
		numericCode:       68,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "boliviano",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BRL"] = CommonCurrency{
		name:              "Brazilian Real",
		symbol:            "R$",
		code:              "BRL",
		numericCode:       986,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "real",
		minorUnit:         "centavo",
		formatPattern:     "{code} {amount}",
	}

	Currencies["BSD"] = CommonCurrency{
		name:              "Bahamian Dollar",
		symbol:            "$",
		code:              "BSD",
		numericCode:       44,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BTN"] = CommonCurrency{
		name:              "Ngultrum",
		symbol:            "Nu.",
		code:              "BTN",
		numericCode:       64,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "ngultrum",
		minorUnit:         "chetrum",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BWP"] = CommonCurrency{
		name:              "Pula",
		symbol:            "P",
		code:              "BWP",
		numericCode:       72,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pula",
		minorUnit:         "thebe",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["BYN"] = CommonCurrency{
		name:              "Belarusian Ruble",
		symbol:            "Br",
		code:              "BYN",
		numericCode:       933,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: " ",
		majorUnit:         "ruble",
		minorUnit:         "kopek",
		formatPattern:     "{amount} {code}",
	}

	Currencies["BZD"] = CommonCurrency{
		name:              "Belize Dollar",
		symbol:            "$",
		code:              "BZD",
		numericCode:       84,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CAD"] = CommonCurrency{
		name:              "Canadian Dollar",
		symbol:            "$",
		code:              "CAD",
		numericCode:       124,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CDF"] = CommonCurrency{
		name:              "Congolese Franc",
		symbol:            "FC",
		code:              "CDF",
		numericCode:       976,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "centime",
		formatPattern:     "{amount} {code}",
	}

	Currencies["CHF"] = CommonCurrency{
		name:              "Swiss Franc",
		symbol:            "CHF",
		code:              "CHF",
		numericCode:       756,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: "'",
		majorUnit:         "franc",
		minorUnit:         "centime",
		formatPattern:     "{symbol} {amount}",
	}

	Currencies["CLP"] = CommonCurrency{
		name:              "Chilean Peso",
		symbol:            "$",
		code:              "CLP",
		numericCode:       152,
		decimalCount:      0,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "peso",
		minorUnit:         "",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CNY"] = CommonCurrency{
		name:              "Yuan Renminbi",
		symbol:            "¥",
		code:              "CNY",
		numericCode:       156,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "yuan",
		minorUnit:         "fen",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["COP"] = CommonCurrency{
		name:              "Colombian Peso",
		symbol:            "$",
		code:              "COP",
		numericCode:       170,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CRC"] = CommonCurrency{
		name:              "Costa Rican Colon",
		symbol:            "₡",
		code:              "CRC",
		numericCode:       188,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "colón",
		minorUnit:         "céntimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CUP"] = CommonCurrency{
		name:              "Cuban Peso",
		symbol:            "$",
		code:              "CUP",
		numericCode:       192,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["CVE"] = CommonCurrency{
		name:              "Cabo Verde Escudo",
		symbol:            "$",
		code:              "CVE",
		numericCode:       132,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "escudo",
		minorUnit:         "centavo",
		formatPattern:     "{code} {amount}",
	}

	Currencies["CZK"] = CommonCurrency{
		name:              "Czech Koruna",
		symbol:            "Kč",
		code:              "CZK",
		numericCode:       203,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: " ",
		majorUnit:         "koruna",
		minorUnit:         "haléř",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["DJF"] = CommonCurrency{
		name:              "Djibouti Franc",
		symbol:            "Fdj",
		code:              "DJF",
		numericCode:       262,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["DKK"] = CommonCurrency{
		name:              "Danish Krone",
		symbol:            "kr",
		code:              "DKK",
		numericCode:       208,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "krone",
		minorUnit:         "øre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["DOP"] = CommonCurrency{
		name:              "Dominican Peso",
		symbol:            "$",
		code:              "DOP",
		numericCode:       214,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["DZD"] = CommonCurrency{
		name:              "Algerian Dinar",
		symbol:            "د.ج",
		code:              "DZD",
		numericCode:       12,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "centime",
		formatPattern:     "{amount} {code}",
	}

	Currencies["EGP"] = CommonCurrency{
		name:              "Egyptian Pound",
		symbol:            "E£",
		code:              "EGP",
		numericCode:       818,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "piastre",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["ERN"] = CommonCurrency{
		name:              "Nakfa",
		symbol:            "Nfk",
		code:              "ERN",
		numericCode:       232,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "nakfa",
		minorUnit:         "cent",
		formatPattern:     "{amount} {code}",
	}

	Currencies["ETB"] = CommonCurrency{
		name:              "Ethiopian Birr",
		symbol:            "Br",
		code:              "ETB",
		numericCode:       230,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "birr",
		minorUnit:         "santim",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["EUR"] = CommonCurrency{
		name:              "Euro",
		symbol:            "€",
		code:              "EUR",
		numericCode:       978,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "euro",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["FJD"] = CommonCurrency{
		name:              "Fiji Dollar",
		symbol:            "$",
		code:              "FJD",
		numericCode:       242,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["FKP"] = CommonCurrency{
		name:              "Falkland Islands Pound",
		symbol:            "£",
		code:              "FKP",
		numericCode:       238,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "penny",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GBP"] = CommonCurrency{
		name:              "Pound Sterling",
		symbol:            "£",
		code:              "GBP",
		numericCode:       826,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "penny",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GEL"] = CommonCurrency{
		name:              "Lari",
		symbol:            "₾",
		code:              "GEL",
		numericCode:       981,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: " ",
		majorUnit:         "lari",
		minorUnit:         "tetri",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["GHS"] = CommonCurrency{
		name:              "Ghana Cedi",
		symbol:            "₵",
		code:              "GHS",
		numericCode:       936,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "cedi",
		minorUnit:         "pesewa",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GIP"] = CommonCurrency{
		name:              "Gibraltar Pound",
		symbol:            "£",
		code:              "GIP",
		numericCode:       292,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "penny",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GMD"] = CommonCurrency{
		name:              "Dalasi",
		symbol:            "D",
		code:              "GMD",
		numericCode:       270,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dalasi",
		minorUnit:         "butut",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GNF"] = CommonCurrency{
		name:              "Guinean Franc",
		symbol:            "FG",
		code:              "GNF",
		numericCode:       324,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["GTQ"] = CommonCurrency{
		name:              "Quetzal",
		symbol:            "Q",
		code:              "GTQ",
		numericCode:       320,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "quetzal",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["GYD"] = CommonCurrency{
		name:              "Guyana Dollar",
		symbol:            "$",
		code:              "GYD",
		numericCode:       328,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["HKD"] = CommonCurrency{
		name:              "Hong Kong Dollar",
		symbol:            "HK$",
		code:              "HKD",
		numericCode:       344,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["HNL"] = CommonCurrency{
		name:              "Lempira",
		symbol:            "L",
		code:              "HNL",
		numericCode:       340,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "lempira",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["HTG"] = CommonCurrency{
		name:              "Gourde",
		symbol:            "G",
		code:              "HTG",
		numericCode:       332,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "gourde",
		minorUnit:         "centime",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["HUF"] = CommonCurrency{
		name:              "Forint",
		symbol:            "Ft",
		code:              "HUF",
		numericCode:       348,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: " ",
		majorUnit:         "forint",
		minorUnit:         "fillér",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["IDR"] = CommonCurrency{
		name:              "Rupiah",
		symbol:            "Rp",
		code:              "IDR",
		numericCode:       360,
		decimalCount:      0,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "rupiah",
		minorUnit:         "",
		formatPattern:     "{code} {amount}",
	}

	Currencies["ILS"] = CommonCurrency{
		name:              "New Israeli Sheqel",
		symbol:            "₪",
		code:              "ILS",
		numericCode:       376,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "shekel",
		minorUnit:         "agora",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["INR"] = CommonCurrency{
		name:              "Indian Rupee",
		symbol:            "₹",
		code:              "INR",
		numericCode:       356,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "paisa",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["IQD"] = CommonCurrency{
		name:              "Iraqi Dinar",
		symbol:            "ع.د",
		code:              "IQD",
		numericCode:       368,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "fils",
		formatPattern:     "{amount} {code}",
	}

	Currencies["IRR"] = CommonCurrency{
		name:              "Iranian Rial",
		symbol:            "﷼",
		code:              "IRR",
		numericCode:       364,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rial",
		minorUnit:         "dinar",
		formatPattern:     "{amount} {code}",
	}

	Currencies["ISK"] = CommonCurrency{
		name:              "Iceland Krona",
		symbol:            "kr",
		code:              "ISK",
		numericCode:       352,
		decimalCount:      0,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "króna",
		minorUnit:         "",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["JMD"] = CommonCurrency{
		name:              "Jamaican Dollar",
		symbol:            "$",
		code:              "JMD",
		numericCode:       388,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["JOD"] = CommonCurrency{
		name:              "Jordanian Dinar",
		symbol:            "JD",
		code:              "JOD",
		numericCode:       400,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "piastre",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["JPY"] = CommonCurrency{
		name:              "Yen",
		symbol:            "¥",
		code:              "JPY",
		numericCode:       392,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "yen",
		minorUnit:         "",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KRW"] = CommonCurrency{
		name:              "Won",
		symbol:            "₩",
		code:              "KRW",
		numericCode:       410,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "won",
		minorUnit:         "",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KES"] = CommonCurrency{
		name:              "Kenyan Shilling",
		symbol:            "KSh",
		code:              "KES",
		numericCode:       404,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "shilling",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KGS"] = CommonCurrency{
		name:              "Som",
		symbol:            "сом",
		code:              "KGS",
		numericCode:       417,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "som",
		minorUnit:         "tyiyn",
		formatPattern:     "{amount} {code}",
	}

	Currencies["KHR"] = CommonCurrency{
		name:              "Riel",
		symbol:            "៛",
		code:              "KHR",
		numericCode:       116,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "riel",
		minorUnit:         "sen",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KMF"] = CommonCurrency{
		name:              "Comorian Franc",
		symbol:            "CF",
		code:              "KMF",
		numericCode:       174,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["KPW"] = CommonCurrency{
		name:              "North Korean Won",
		symbol:            "₩",
		code:              "KPW",
		numericCode:       408,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "won",
		minorUnit:         "chon",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KWD"] = CommonCurrency{
		name:              "Kuwaiti Dinar",
		symbol:            "KD",
		code:              "KWD",
		numericCode:       414,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "fils",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["KYD"] = CommonCurrency{
		name:              "Cayman Islands Dollar",
		symbol:            "$",
		code:              "KYD",
		numericCode:       136,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["LAK"] = CommonCurrency{
		name:              "Lao Kip",
		symbol:            "₭",
		code:              "LAK",
		numericCode:       418,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kip",
		minorUnit:         "att",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["LBP"] = CommonCurrency{
		name:              "Lebanese Pound",
		symbol:            "ل.ل",
		code:              "LBP",
		numericCode:       422,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "piastre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["LKR"] = CommonCurrency{
		name:              "Sri Lanka Rupee",
		symbol:            "Rs",
		code:              "LKR",
		numericCode:       144,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["LRD"] = CommonCurrency{
		name:              "Liberian Dollar",
		symbol:            "$",
		code:              "LRD",
		numericCode:       430,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["LSL"] = CommonCurrency{
		name:              "Loti",
		symbol:            "L",
		code:              "LSL",
		numericCode:       426,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "loti",
		minorUnit:         "sente",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["LYD"] = CommonCurrency{
		name:              "Libyan Dinar",
		symbol:            "LD",
		code:              "LYD",
		numericCode:       434,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "dirham",
		formatPattern:     "{amount} {code}",
	}

	Currencies["MAD"] = CommonCurrency{
		name:              "Moroccan Dirham",
		symbol:            "MAD",
		code:              "MAD",
		numericCode:       504,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dirham",
		minorUnit:         "centime",
		formatPattern:     "{amount} {code}",
	}

	Currencies["MDL"] = CommonCurrency{
		name:              "Moldovan Leu",
		symbol:            "L",
		code:              "MDL",
		numericCode:       498,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "leu",
		minorUnit:         "ban",
		formatPattern:     "{amount} {code}",
	}

	Currencies["MGA"] = CommonCurrency{
		name:              "Malagasy Ariary",
		symbol:            "Ar",
		code:              "MGA",
		numericCode:       969,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "ariary",
		minorUnit:         "iraimbilanja",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MKD"] = CommonCurrency{
		name:              "Denar",
		symbol:            "ден",
		code:              "MKD",
		numericCode:       807,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "denar",
		minorUnit:         "deni",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["MMK"] = CommonCurrency{
		name:              "Kyat",
		symbol:            "K",
		code:              "MMK",
		numericCode:       104,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kyat",
		minorUnit:         "pya",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MNT"] = CommonCurrency{
		name:              "Tugrik",
		symbol:            "₮",
		code:              "MNT",
		numericCode:       496,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "tugrik",
		minorUnit:         "möngö",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MOP"] = CommonCurrency{
		name:              "Pataca",
		symbol:            "MOP$",
		code:              "MOP",
		numericCode:       446,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pataca",
		minorUnit:         "avo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MRU"] = CommonCurrency{
		name:              "Ouguiya",
		symbol:            "UM",
		code:              "MRU",
		numericCode:       929,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "ouguiya",
		minorUnit:         "khoum",
		formatPattern:     "{amount} {code}",
	}

	Currencies["MUR"] = CommonCurrency{
		name:              "Mauritius Rupee",
		symbol:            "Rs",
		code:              "MUR",
		numericCode:       480,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MVR"] = CommonCurrency{
		name:              "Rufiyaa",
		symbol:            "Rf",
		code:              "MVR",
		numericCode:       462,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rufiyaa",
		minorUnit:         "laari",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MWK"] = CommonCurrency{
		name:              "Malawi Kwacha",
		symbol:            "MK",
		code:              "MWK",
		numericCode:       454,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kwacha",
		minorUnit:         "tambala",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MXN"] = CommonCurrency{
		name:              "Mexican Peso",
		symbol:            "$",
		code:              "MXN",
		numericCode:       484,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MYR"] = CommonCurrency{
		name:              "Malaysian Ringgit",
		symbol:            "RM",
		code:              "MYR",
		numericCode:       458,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "ringgit",
		minorUnit:         "sen",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["MZN"] = CommonCurrency{
		name:              "Mozambique Metical",
		symbol:            "MT",
		code:              "MZN",
		numericCode:       943,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "metical",
		minorUnit:         "centavo",
		formatPattern:     "{amount} {code}",
	}

	Currencies["NAD"] = CommonCurrency{
		name:              "Namibia Dollar",
		symbol:            "$",
		code:              "NAD",
		numericCode:       516,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["NGN"] = CommonCurrency{
		name:              "Naira",
		symbol:            "₦",
		code:              "NGN",
		numericCode:       566,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "naira",
		minorUnit:         "kobo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["NIO"] = CommonCurrency{
		name:              "Cordoba Oro",
		symbol:            "C$",
		code:              "NIO",
		numericCode:       558,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "córdoba",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["NOK"] = CommonCurrency{
		name:              "Norwegian Krone",
		symbol:            "kr",
		code:              "NOK",
		numericCode:       578,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "krone",
		minorUnit:         "øre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["NPR"] = CommonCurrency{
		name:              "Nepalese Rupee",
		symbol:            "Rs",
		code:              "NPR",
		numericCode:       524,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "paisa",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["NZD"] = CommonCurrency{
		name:              "New Zealand Dollar",
		symbol:            "$",
		code:              "NZD",
		numericCode:       554,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["OMR"] = CommonCurrency{
		name:              "Rial Omani",
		symbol:            "ر.ع.",
		code:              "OMR",
		numericCode:       512,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rial",
		minorUnit:         "baisa",
		formatPattern:     "{amount} {code}",
	}

	Currencies["PAB"] = CommonCurrency{
		name:              "Balboa",
		symbol:            "B/.",
		code:              "PAB",
		numericCode:       590,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "balboa",
		minorUnit:         "centésimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["PEN"] = CommonCurrency{
		name:              "Sol",
		symbol:            "S/",
		code:              "PEN",
		numericCode:       604,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "sol",
		minorUnit:         "céntimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["PGK"] = CommonCurrency{
		name:              "Kina",
		symbol:            "K",
		code:              "PGK",
		numericCode:       598,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kina",
		minorUnit:         "toea",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["PHP"] = CommonCurrency{
		name:              "Philippine Peso",
		symbol:            "₱",
		code:              "PHP",
		numericCode:       608,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "peso",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["PKR"] = CommonCurrency{
		name:              "Pakistan Rupee",
		symbol:            "Rs",
		code:              "PKR",
		numericCode:       586,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "paisa",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["PLN"] = CommonCurrency{
		name:              "Zloty",
		symbol:            "zł",
		code:              "PLN",
		numericCode:       985,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "zloty",
		minorUnit:         "grosz",
		formatPattern:     "{amount} {code}",
	}

	Currencies["PYG"] = CommonCurrency{
		name:              "Guarani",
		symbol:            "₲",
		code:              "PYG",
		numericCode:       600,
		decimalCount:      0,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "guaraní",
		minorUnit:         "",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["QAR"] = CommonCurrency{
		name:              "Qatari Rial",
		symbol:            "QR",
		code:              "QAR",
		numericCode:       634,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rial",
		minorUnit:         "dirham",
		formatPattern:     "{code} {amount}",
	}

	Currencies["RON"] = CommonCurrency{
		name:              "Romanian Leu",
		symbol:            "lei",
		code:              "RON",
		numericCode:       946,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "leu",
		minorUnit:         "ban",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["RSD"] = CommonCurrency{
		name:              "Serbian Dinar",
		symbol:            "din",
		code:              "RSD",
		numericCode:       941,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "dinar",
		minorUnit:         "para",
		formatPattern:     "{amount} {code}",
	}

	Currencies["RUB"] = CommonCurrency{
		name:              "Russian Ruble",
		symbol:            "₽",
		code:              "RUB",
		numericCode:       643,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "ruble",
		minorUnit:         "kopek",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["RWF"] = CommonCurrency{
		name:              "Rwanda Franc",
		symbol:            "RF",
		code:              "RWF",
		numericCode:       646,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["SAR"] = CommonCurrency{
		name:              "Saudi Riyal",
		symbol:            "ر.س",
		code:              "SAR",
		numericCode:       682,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "riyal",
		minorUnit:         "halala",
		formatPattern:     "{code} {amount}",
	}

	Currencies["SBD"] = CommonCurrency{
		name:              "Solomon Islands Dollar",
		symbol:            "$",
		code:              "SBD",
		numericCode:       90,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SCR"] = CommonCurrency{
		name:              "Seychelles Rupee",
		symbol:            "Rs",
		code:              "SCR",
		numericCode:       690,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rupee",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SDG"] = CommonCurrency{
		name:              "Sudanese Pound",
		symbol:            "£",
		code:              "SDG",
		numericCode:       938,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "piastre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["SEK"] = CommonCurrency{
		name:              "Swedish Krona",
		symbol:            "kr",
		code:              "SEK",
		numericCode:       752,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "krona",
		minorUnit:         "öre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["SGD"] = CommonCurrency{
		name:              "Singapore Dollar",
		symbol:            "$",
		code:              "SGD",
		numericCode:       702,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SHP"] = CommonCurrency{
		name:              "Saint Helena Pound",
		symbol:            "£",
		code:              "SHP",
		numericCode:       654,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "penny",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SLE"] = CommonCurrency{
		name:              "Leone",
		symbol:            "Le",
		code:              "SLE",
		numericCode:       925,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "leone",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SOS"] = CommonCurrency{
		name:              "Somali Shilling",
		symbol:            "Sh",
		code:              "SOS",
		numericCode:       706,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "shilling",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SRD"] = CommonCurrency{
		name:              "Surinam Dollar",
		symbol:            "$",
		code:              "SRD",
		numericCode:       968,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SSP"] = CommonCurrency{
		name:              "South Sudanese Pound",
		symbol:            "£",
		code:              "SSP",
		numericCode:       728,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "piastre",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["STN"] = CommonCurrency{
		name:              "Dobra",
		symbol:            "Db",
		code:              "STN",
		numericCode:       930,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dobra",
		minorUnit:         "cêntimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SVC"] = CommonCurrency{
		name:              "El Salvador Colon",
		symbol:            "₡",
		code:              "SVC",
		numericCode:       222,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "colón",
		minorUnit:         "centavo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["SYP"] = CommonCurrency{
		name:              "Syrian Pound",
		symbol:            "£",
		code:              "SYP",
		numericCode:       760,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pound",
		minorUnit:         "piastre",
		formatPattern:     "{amount} {code}",
	}

	Currencies["SZL"] = CommonCurrency{
		name:              "Lilangeni",
		symbol:            "E",
		code:              "SZL",
		numericCode:       748,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "lilangeni",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["THB"] = CommonCurrency{
		name:              "Baht",
		symbol:            "฿",
		code:              "THB",
		numericCode:       764,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "baht",
		minorUnit:         "satang",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["TJS"] = CommonCurrency{
		name:              "Somoni",
		symbol:            "SM",
		code:              "TJS",
		numericCode:       972,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "somoni",
		minorUnit:         "diram",
		formatPattern:     "{amount} {code}",
	}

	Currencies["TMT"] = CommonCurrency{
		name:              "Turkmenistan New Manat",
		symbol:            "T",
		code:              "TMT",
		numericCode:       934,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "manat",
		minorUnit:         "tennesi",
		formatPattern:     "{amount} {code}",
	}

	Currencies["TND"] = CommonCurrency{
		name:              "Tunisian Dinar",
		symbol:            "DT",
		code:              "TND",
		numericCode:       788,
		decimalCount:      3,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dinar",
		minorUnit:         "millime",
		formatPattern:     "{amount} {code}",
	}

	Currencies["TOP"] = CommonCurrency{
		name:              "Pa'anga",
		symbol:            "T$",
		code:              "TOP",
		numericCode:       776,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "pa'anga",
		minorUnit:         "seniti",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["TRY"] = CommonCurrency{
		name:              "Turkish Lira",
		symbol:            "₺",
		code:              "TRY",
		numericCode:       949,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "lira",
		minorUnit:         "kuruş",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["TTD"] = CommonCurrency{
		name:              "Trinidad and Tobago Dollar",
		symbol:            "$",
		code:              "TTD",
		numericCode:       780,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["TWD"] = CommonCurrency{
		name:              "New Taiwan Dollar",
		symbol:            "NT$",
		code:              "TWD",
		numericCode:       901,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["TZS"] = CommonCurrency{
		name:              "Tanzanian Shilling",
		symbol:            "TSh",
		code:              "TZS",
		numericCode:       834,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "shilling",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["UAH"] = CommonCurrency{
		name:              "Hryvnia",
		symbol:            "₴",
		code:              "UAH",
		numericCode:       980,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: " ",
		majorUnit:         "hryvnia",
		minorUnit:         "kopiyka",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["UGX"] = CommonCurrency{
		name:              "Uganda Shilling",
		symbol:            "USh",
		code:              "UGX",
		numericCode:       800,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "shilling",
		minorUnit:         "",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["USD"] = CommonCurrency{
		name:              "US Dollar",
		symbol:            "$",
		code:              "USD",
		numericCode:       840,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["UYU"] = CommonCurrency{
		name:              "Peso Uruguayo",
		symbol:            "$U",
		code:              "UYU",
		numericCode:       858,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "peso",
		minorUnit:         "centésimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["UZS"] = CommonCurrency{
		name:              "Uzbekistan Sum",
		symbol:            "сўм",
		code:              "UZS",
		numericCode:       860,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "sum",
		minorUnit:         "tiyin",
		formatPattern:     "{amount} {code}",
	}

	Currencies["VES"] = CommonCurrency{
		name:              "Bolívar Soberano",
		symbol:            "Bs.S",
		code:              "VES",
		numericCode:       928,
		decimalCount:      2,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "bolívar",
		minorUnit:         "céntimo",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["VND"] = CommonCurrency{
		name:              "Dong",
		symbol:            "₫",
		code:              "VND",
		numericCode:       704,
		decimalCount:      0,
		decimalSeparator:  ",",
		thousandSeparator: ".",
		majorUnit:         "dong",
		minorUnit:         "",
		formatPattern:     "{amount} {symbol}",
	}

	Currencies["VUV"] = CommonCurrency{
		name:              "Vatu",
		symbol:            "VT",
		code:              "VUV",
		numericCode:       548,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "vatu",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["WST"] = CommonCurrency{
		name:              "Tala",
		symbol:            "T",
		code:              "WST",
		numericCode:       882,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "tala",
		minorUnit:         "sene",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["XAF"] = CommonCurrency{
		name:              "CFA Franc BEAC",
		symbol:            "FCFA",
		code:              "XAF",
		numericCode:       950,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["XCD"] = CommonCurrency{
		name:              "East Caribbean Dollar",
		symbol:            "$",
		code:              "XCD",
		numericCode:       951,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "dollar",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["XCG"] = CommonCurrency{
		name:              "Caribbean Guilder",
		symbol:            "ƒ",
		code:              "XCG",
		numericCode:       532,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "guilder",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["XOF"] = CommonCurrency{
		name:              "CFA Franc BCEAO",
		symbol:            "CFA",
		code:              "XOF",
		numericCode:       952,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["XPF"] = CommonCurrency{
		name:              "CFP Franc",
		symbol:            "F",
		code:              "XPF",
		numericCode:       953,
		decimalCount:      0,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "franc",
		minorUnit:         "",
		formatPattern:     "{amount} {code}",
	}

	Currencies["YER"] = CommonCurrency{
		name:              "Yemeni Rial",
		symbol:            "﷼",
		code:              "YER",
		numericCode:       886,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rial",
		minorUnit:         "fils",
		formatPattern:     "{amount} {code}",
	}

	Currencies["ZAR"] = CommonCurrency{
		name:              "Rand",
		symbol:            "R",
		code:              "ZAR",
		numericCode:       710,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "rand",
		minorUnit:         "cent",
		formatPattern:     "{code} {amount}",
	}

	Currencies["ZMW"] = CommonCurrency{
		name:              "Zambian Kwacha",
		symbol:            "ZK",
		code:              "ZMW",
		numericCode:       967,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "kwacha",
		minorUnit:         "ngwee",
		formatPattern:     "{symbol}{amount}",
	}

	Currencies["ZWG"] = CommonCurrency{
		name:              "Zimbabwe Gold",
		symbol:            "ZiG",
		code:              "ZWG",
		numericCode:       924,
		decimalCount:      2,
		decimalSeparator:  ".",
		thousandSeparator: ",",
		majorUnit:         "ZiG",
		minorUnit:         "cent",
		formatPattern:     "{symbol}{amount}",
	}
}
