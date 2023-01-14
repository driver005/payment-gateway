package ach

import (
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

// converters handles golang to ACH type Converters
type converters struct{}

func (c *converters) parseNumField(r string) (s int) {
	s, _ = strconv.Atoi(strings.TrimSpace(r))
	return s
}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

// formatSimpleDate takes a YYMMDD date and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleDate(s string) string {
	if s == "" {
		return c.stringField(s, 6)
	}
	return s
}

// formatSimpleTime takes a HHmm (H=hour, m=minute) time and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleTime(s string) string {
	if s == "" {
		return c.stringField(s, 4)
	}
	return s
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(utf8.RuneCountInString(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}

// numericField right-justified, unsigned, and zero filled
func (c *converters) numericField(n int, max uint) string {
	s := strconv.Itoa(n)
	ln := uint(utf8.RuneCountInString(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

// stringField slices to max length and zero filled
func (c *converters) stringField(s string, max uint) string {
	ln := uint(utf8.RuneCountInString(s))
	if ln > max {
		return s[:max]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

// leastSignificantDigits returns the least significant digits of v limited by maxDigits.
func (c *converters) leastSignificantDigits(v int, maxDigits uint) int {
	return v % int(math.Pow10(int(maxDigits)))
}
