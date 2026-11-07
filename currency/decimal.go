package currency

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	// ErrInvalidDecimal is returned for a string that is not a plain decimal number.
	ErrInvalidDecimal = errors.New("invalid decimal amount")
	// ErrDecimalPrecision is returned for a decimal with more significant decimal places
	// than the currency has.
	ErrDecimalPrecision = errors.New("decimal amount is more precise than the currency")
	// ErrDecimalOverflow is returned for a decimal too large for an int64 of minor units.
	ErrDecimalOverflow = errors.New("decimal amount overflows minor units")
)

func (c CommonCurrency) FormatDecimal(amount int64) string {
	negative := amount < 0
	// Kept as uint64 so the magnitude of math.MinInt64 does not overflow.
	magnitude := uint64(amount)
	if negative {
		magnitude = uint64(-(amount + 1)) + 1
	}

	digits := strconv.FormatUint(magnitude, 10)
	if places := int(c.decimalCount); places > 0 {
		if len(digits) <= places {
			digits = strings.Repeat("0", places-len(digits)+1) + digits
		}
		digits = digits[:len(digits)-places] + "." + digits[len(digits)-places:]
	}
	if negative {
		return "-" + digits
	}
	return digits
}

func (c CommonCurrency) ParseDecimal(amount string) (int64, error) {
	s := strings.TrimSpace(amount)
	negative := false
	if s != "" && (s[0] == '-' || s[0] == '+') {
		negative = s[0] == '-'
		s = s[1:]
	}

	whole, fraction, _ := strings.Cut(s, ".")
	if whole == "" && fraction == "" || !allDigits(whole) || !allDigits(fraction) {
		return 0, ErrInvalidDecimal
	}

	places := int(c.decimalCount)
	if len(fraction) > places {
		if strings.Trim(fraction[places:], "0") != "" {
			return 0, ErrDecimalPrecision
		}
		fraction = fraction[:places]
	}
	fraction += strings.Repeat("0", places-len(fraction))

	var units uint64
	for _, d := range whole + fraction {
		digit := uint64(d - '0')
		if units > (math.MaxInt64-digit)/10 {
			return 0, ErrDecimalOverflow
		}
		units = units*10 + digit
	}
	if negative {
		return -int64(units), nil
	}
	return int64(units), nil
}

func allDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
