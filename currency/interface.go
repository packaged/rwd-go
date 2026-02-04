package currency

type Currency interface {
	getSymbol() string
	getCode() string
	getNumericCode() uint16
	getName() string
	getDecimalCount() uint8
	getDecimalSeparator() string
	getThousandSeparator() string
	getMajorUnit() string
	getMinorUnit() string
	getFormat() string
	Format(amount int64) string
	format(amount int64, pattern string) string
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

func (c CommonCurrency) getSymbol() string {
	return c.symbol
}

func (c CommonCurrency) getCode() string {
	return c.code
}

func (c CommonCurrency) getNumericCode() uint16 {
	return c.numericCode
}

func (c CommonCurrency) getName() string {
	return c.name
}

func (c CommonCurrency) getDecimalCount() uint8 {
	return c.decimalCount
}

func (c CommonCurrency) getDecimalSeparator() string {
	return c.decimalSeparator
}

func (c CommonCurrency) getThousandSeparator() string {
	return c.thousandSeparator
}

func (c CommonCurrency) getMajorUnit() string {
	return c.majorUnit
}

func (c CommonCurrency) getMinorUnit() string {
	return c.minorUnit
}

func (c CommonCurrency) getFormat() string {
	return c.formatPattern
}

// Format formats the amount using the currency's default format pattern
func (c CommonCurrency) Format(amount int64) string {
	return c.format(amount, c.formatPattern)
}

func (c CommonCurrency) format(amount int64, pattern string) string {
	// Format the amount according to currency rules
	formattedAmount := c.formatAmount(amount)

	// Replace placeholders in pattern
	result := pattern
	result = replaceAll(result, "{symbol}", c.symbol)
	result = replaceAll(result, "{code}", c.code)
	result = replaceAll(result, "{amount}", formattedAmount)

	return result
}

// formatAmount formats the amount with proper decimal and thousand separators
func (c CommonCurrency) formatAmount(amount int64) string {
	// Handle negative amounts
	negative := amount < 0
	if negative {
		amount = -amount
	}

	var intPart, fracPart int64
	if c.decimalCount > 0 {
		divisor := int64(1)
		for i := uint8(0); i < c.decimalCount; i++ {
			divisor *= 10
		}
		intPart = amount / divisor
		fracPart = amount % divisor
	} else {
		intPart = amount
		fracPart = 0
	}

	// Format integer part with thousand separators
	intStr := formatWithSeparator(intPart, c.thousandSeparator)

	// Build result
	var result string
	if negative {
		result = "-"
	}

	result += intStr

	// Add decimal part if needed
	if c.decimalCount > 0 {
		result += c.decimalSeparator
		// Pad fractional part with leading zeros
		fracStr := ""
		divisor := int64(1)
		for i := uint8(0); i < c.decimalCount; i++ {
			divisor *= 10
		}
		divisor /= 10
		for i := uint8(0); i < c.decimalCount; i++ {
			digit := (fracPart / divisor) % 10
			fracStr += string('0' + byte(digit))
			divisor /= 10
		}
		result += fracStr
	}

	return result
}

// formatWithSeparator adds thousand separators to a number
func formatWithSeparator(n int64, separator string) string {
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
