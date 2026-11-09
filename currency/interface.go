package currency

import (
	"strconv"
	"strings"
)

type Currency interface {
	GetSymbol() string
	GetCode() string
	GetNumericCode() uint16
	GetName() string
	GetDecimalCount() uint8
	GetDecimalSeparator() string
	GetThousandSeparator() string
	GetMajorUnit() string
	GetMinorUnit() string
	GetFormat() string
	Format(amount int64) string
	// CustomFormat formats the given amount into a currency string based on the specified pattern available tokens:
	// {symbol}, {code}, {amount}
	CustomFormat(amount int64, pattern string) string
	FromInt(amount int64) float64
	FromInt32(amount int32) float64
	// FormatDecimal formats an amount in minor units as a plain decimal string with the
	// currency's decimal places, e.g. 2999 USD is "29.99", 5 USD is "0.05", 1000 JPY is
	// "1000". It uses "." and no thousand separators, for machine use.
	FormatDecimal(amount int64) string
	// ParseDecimal parses a plain decimal string into minor units, e.g. "29.99" USD is
	// 2999. It rejects more decimal places than the currency has, unless they are zeros.
	ParseDecimal(amount string) (int64, error)
}

type CommonCurrency struct {
	name              string
	symbol            string
	code              string
	numericCode       uint16
	decimalCount      uint8
	decimalSeparator  string
	thousandSeparator string
	majorUnit         string
	minorUnit         string
	formatPattern     string
}

func (c CommonCurrency) GetSymbol() string {
	return c.symbol
}

func (c CommonCurrency) GetCode() string {
	return c.code
}

func (c CommonCurrency) GetNumericCode() uint16 {
	return c.numericCode
}

func (c CommonCurrency) GetName() string {
	return c.name
}

func (c CommonCurrency) GetDecimalCount() uint8 {
	return c.decimalCount
}

func (c CommonCurrency) GetDecimalSeparator() string {
	return c.decimalSeparator
}

func (c CommonCurrency) GetThousandSeparator() string {
	return c.thousandSeparator
}

func (c CommonCurrency) GetMajorUnit() string {
	return c.majorUnit
}

func (c CommonCurrency) GetMinorUnit() string {
	return c.minorUnit
}

func (c CommonCurrency) GetFormat() string {
	return c.formatPattern
}

// Format formats the amount using the currency's default format pattern
func (c CommonCurrency) Format(amount int64) string {
	return c.CustomFormat(amount, c.formatPattern)
}

// CustomFormat formats the amount in minor units with pattern, replacing {symbol},
// {code} and {amount}. A negative amount's minus sign leads the result ("-$0.50",
// "-0,50 €") when the pattern has an {amount}; a pattern without one is used as written.
func (c CommonCurrency) CustomFormat(amount int64, pattern string) string {
	result := pattern
	result = replaceAll(result, "{symbol}", c.symbol)
	result = replaceAll(result, "{code}", c.code)
	if !strings.Contains(result, "{amount}") {
		return result
	}
	result = replaceAll(result, "{amount}", c.formatAmount(magnitude(amount)))
	if amount < 0 {
		result = "-" + result
	}
	return result
}

// magnitude returns the absolute value of amount, including math.MinInt64.
func magnitude(amount int64) uint64 {
	if amount < 0 {
		return uint64(-(amount + 1)) + 1
	}
	return uint64(amount)
}

// FromInt converts a minor unit int64 amount to a float64 major unit amount
// e.g. for USD (2 decimals): FromInt(12345) returns 123.45
func (c CommonCurrency) FromInt(amount int64) float64 {
	if c.decimalCount == 0 {
		return float64(amount)
	}
	divisor := int64(1)
	for i := uint8(0); i < c.decimalCount; i++ {
		divisor *= 10
	}
	return float64(amount) / float64(divisor)
}

// FromInt32 converts a minor unit int32 amount to a float64 major unit amount
func (c CommonCurrency) FromInt32(amount int32) float64 {
	return c.FromInt(int64(amount))
}

// formatAmount formats a non-negative amount in minor units with the currency's decimal
// and thousand separators.
func (c CommonCurrency) formatAmount(amount uint64) string {
	divisor := uint64(1)
	for i := uint8(0); i < c.decimalCount; i++ {
		divisor *= 10
	}
	result := formatWithSeparator(amount/divisor, c.thousandSeparator)
	if c.decimalCount > 0 {
		frac := strconv.FormatUint(amount%divisor, 10)
		result += c.decimalSeparator + strings.Repeat("0", int(c.decimalCount)-len(frac)) + frac
	}
	return result
}

// formatWithSeparator adds thousand separators to a number
func formatWithSeparator(n uint64, separator string) string {
	if n == 0 {
		return "0"
	}

	// Convert to string in reverse
	digits := []byte{}
	for n > 0 {
		digits = append(digits, byte('0'+n%10))
		n /= 10
	}

	// Build result with separators (from right to left, add separator every 3 digits)
	result := ""
	for i := 0; i < len(digits); i++ {
		if i > 0 && i%3 == 0 {
			result = separator + result
		}
		result = string(digits[i]) + result
	}

	return result
}

// replaceAll replaces all occurrences of old with new in s
func replaceAll(s, old, new string) string {
	result := ""
	for len(s) > 0 {
		index := -1
		for i := 0; i <= len(s)-len(old); i++ {
			match := true
			for j := 0; j < len(old); j++ {
				if s[i+j] != old[j] {
					match = false
					break
				}
			}
			if match {
				index = i
				break
			}
		}

		if index == -1 {
			result += s
			break
		}

		result += s[:index] + new
		s = s[index+len(old):]
	}

	return result
}
