package country

type Country interface {
	GetName() string
	GetIso2() string
	GetIso3() string
	GetNumericCode() uint16
	GetDialPrefix() uint16
	GetCurrencyCode() string
}

type CommonCountry struct {
	name         string
	iso2         string
	iso3         string
	numericCode  uint16
	dialPrefix   uint16
	currencyCode string
}

func (c CommonCountry) GetName() string {
	return c.name
}

func (c CommonCountry) GetIso2() string {
	return c.iso2
}

func (c CommonCountry) GetIso3() string {
	return c.iso3
}

func (c CommonCountry) GetNumericCode() uint16 {
	return c.numericCode
}

func (c CommonCountry) GetDialPrefix() uint16 {
	return c.dialPrefix
}

func (c CommonCountry) GetCurrencyCode() string {
	return c.currencyCode
}
