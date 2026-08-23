package money

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Amount is a quantity of a single currency, stored as an integer count
// of that currency's minor units. There is no float anywhere in this
// package: every operation on an Amount is exact.
type Amount struct {
	units    int64
	currency Currency
}

// New builds an Amount directly from a minor-unit count, e.g.
// New(1050, USD) is $10.50.
func New(units int64, c Currency) Amount {
	return Amount{units: units, currency: c}
}

// Zero returns a zero-value Amount in the given currency.
func Zero(c Currency) Amount {
	return Amount{currency: c}
}

func (a Amount) Units() int64 {
	return a.units
}

func (a Amount) Currency() Currency {
	return a.currency
}

// Parse reads a decimal string such as "12.34", "-5", or "1234" into an
// Amount of the given currency. The number of digits after the decimal
// point must not exceed the currency's minor-unit precision; Parse does
// not round, it rejects amounts it cannot represent exactly.
func Parse(s string, c Currency) (Amount, error) {
	orig := s
	s = strings.TrimSpace(s)
	if s == "" {
		return Amount{}, fmt.Errorf("money: cannot parse empty string as %s amount", c.Code)
	}

	neg := false
	switch s[0] {
	case '-':
		neg = true
		s = s[1:]
	case '+':
		s = s[1:]
	}

	intPart, fracPart, hasFrac := strings.Cut(s, ".")
	if intPart == "" || (hasFrac && fracPart == "") {
		return Amount{}, fmt.Errorf("money: invalid amount %q", orig)
	}
	if !isDigits(intPart) || !isDigits(fracPart) {
		return Amount{}, fmt.Errorf("money: invalid amount %q", orig)
	}
	if len(fracPart) > c.Minor {
		return Amount{}, fmt.Errorf("money: %q has more precision than %s supports (%d decimal places)", orig, c.Code, c.Minor)
	}
	fracPart += strings.Repeat("0", c.Minor-len(fracPart))

	units, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Amount{}, fmt.Errorf("money: invalid amount %q: %w", orig, err)
	}
	if neg {
		units = -units
	}
	return Amount{units: units, currency: c}, nil
}

func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// String formats the amount with its minor units placed after a decimal
// point, e.g. "10.50 USD" or "1234 JPY".
func (a Amount) String() string {
	sign := ""
	units := a.units
	if units < 0 {
		sign = "-"
		units = -units
	}
	if a.currency.Minor == 0 {
		return fmt.Sprintf("%s%d %s", sign, units, a.currency.Code)
	}
	scale := pow10(a.currency.Minor)
	whole := units / scale
	frac := units % scale
	return fmt.Sprintf("%s%d.%0*d %s", sign, whole, a.currency.Minor, frac, a.currency.Code)
}

func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

// ErrCurrencyMismatch is returned by operations that require two
// amounts to share a currency when they do not.
var ErrCurrencyMismatch = errors.New("money: currency mismatch")

// Add returns a + b. Both amounts must share a currency.
func (a Amount) Add(b Amount) (Amount, error) {
	if a.currency != b.currency {
		return Amount{}, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, a.currency, b.currency)
	}
	sum := a.units + b.units
	if (b.units > 0 && sum < a.units) || (b.units < 0 && sum > a.units) {
		return Amount{}, fmt.Errorf("money: overflow adding %s and %s", a, b)
	}
	return Amount{units: sum, currency: a.currency}, nil
}

// Sub returns a - b. Both amounts must share a currency.
func (a Amount) Sub(b Amount) (Amount, error) {
	if a.currency != b.currency {
		return Amount{}, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, a.currency, b.currency)
	}
	diff := a.units - b.units
	if (b.units < 0 && diff < a.units) || (b.units > 0 && diff > a.units) {
		return Amount{}, fmt.Errorf("money: overflow subtracting %s from %s", b, a)
	}
	return Amount{units: diff, currency: a.currency}, nil
}

// Cmp compares two amounts of the same currency, returning -1, 0, or 1.
func (a Amount) Cmp(b Amount) (int, error) {
	if a.currency != b.currency {
		return 0, fmt.Errorf("%w: %s and %s", ErrCurrencyMismatch, a.currency, b.currency)
	}
	switch {
	case a.units < b.units:
		return -1, nil
	case a.units > b.units:
		return 1, nil
	default:
		return 0, nil
	}
}
