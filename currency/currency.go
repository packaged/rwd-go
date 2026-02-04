package currency

var Currencies map[string]Currency

func init() {
	Currencies = make(map[string]Currency)

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
}
