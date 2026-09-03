package money

import (
	"errors"
	"math"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		currency Currency
		want     int64
	}{
		{"plain decimal", "12.34", USD, 1234},
		{"negative", "-5", USD, -500},
		{"explicit plus sign", "+3.50", USD, 350},
		{"no fractional part", "1234", JPY, 1234},
		{"surrounding whitespace", "  10.00  ", USD, 1000},
		{"zero", "0", USD, 0},
		{"negative fraction below one unit", "-0.01", USD, -1},
		{"three decimal place currency", "1.234", KWD, 1234},
		{"fewer digits than currency precision", "1.2", KWD, 1200},
		{"negative zero collapses to zero", "-0", USD, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.input, tc.currency)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", tc.input, err)
			}
			if got.Units() != tc.want {
				t.Errorf("Parse(%q) = %d units, want %d", tc.input, got.Units(), tc.want)
			}
			if got.Currency() != tc.currency {
				t.Errorf("Parse(%q) currency = %s, want %s", tc.input, got.Currency(), tc.currency)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		currency Currency
	}{
		{"empty string", "", USD},
		{"only whitespace", "   ", USD},
		{"just a decimal point", ".", USD},
		{"missing integer part", ".5", USD},
		{"missing fractional part", "5.", USD},
		{"more precision than currency supports", "12.345", USD},
		{"more precision than a zero-decimal currency supports", "12.3", JPY},
		{"letters in the integer part", "ab.3", USD},
		{"letters in the fractional part", "12.3a", USD},
		{"stray sign inside the number", "12-3", USD},
		{"double sign", "--5", USD},
		{"two decimal points", "1.2.3", USD},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := Parse(tc.input, tc.currency); err == nil {
				t.Fatalf("Parse(%q) expected an error, got nil", tc.input)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	sum, err := New(150, USD).Add(New(250, USD))
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}
	if sum.Units() != 400 {
		t.Errorf("sum = %d, want 400", sum.Units())
	}

	if _, err := New(100, USD).Add(New(100, EUR)); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("Add across currencies: got %v, want ErrCurrencyMismatch", err)
	}

	if _, err := New(math.MaxInt64, USD).Add(New(1, USD)); err == nil {
		t.Fatalf("Add overflow: expected an error, got nil")
	}

	if _, err := New(math.MinInt64, USD).Add(New(-1, USD)); err == nil {
		t.Fatalf("Add underflow: expected an error, got nil")
	}
}

func TestSub(t *testing.T) {
	diff, err := New(250, USD).Sub(New(150, USD))
	if err != nil {
		t.Fatalf("Sub returned error: %v", err)
	}
	if diff.Units() != 100 {
		t.Errorf("diff = %d, want 100", diff.Units())
	}

	negative, err := New(100, USD).Sub(New(250, USD))
	if err != nil {
		t.Fatalf("Sub returned error: %v", err)
	}
	if negative.Units() != -150 {
		t.Errorf("diff = %d, want -150", negative.Units())
	}

	if _, err := New(100, USD).Sub(New(100, EUR)); !errors.Is(err, ErrCurrencyMismatch) {
		t.Errorf("Sub across currencies: got %v, want ErrCurrencyMismatch", err)
	}

	if _, err := New(math.MinInt64, USD).Sub(New(1, USD)); err == nil {
		t.Fatalf("Sub underflow: expected an error, got nil")
	}

	if _, err := New(math.MaxInt64, USD).Sub(New(-1, USD)); err == nil {
		t.Fatalf("Sub overflow: expected an error, got nil")
	}
}
