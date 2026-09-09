package money

import (
	"math"
	"testing"
)

func TestMul(t *testing.T) {
	cases := []struct {
		name       string
		amount     Amount
		num, denom int64
		mode       RoundingMode
		want       int64
	}{
		{"8.25% tax on 19.99, half up", New(1999, USD), 825, 10000, RoundHalfUp, 165},
		{"15% discount taken off 19.99, half up", New(1999, USD), 85, 100, RoundHalfUp, 1699},
		{"exact quarter needs no rounding", New(100, USD), 1, 4, RoundHalfUp, 25},
		{"halfway rounds up away from zero, half up", New(5, USD), 1, 2, RoundHalfUp, 3},
		{"halfway rounds to even quotient, half even", New(5, USD), 1, 2, RoundHalfEven, 2},
		{"halfway with an already-even quotient stays put, half even", New(1, USD), 1, 2, RoundHalfEven, 0},
		{"negative halfway rounds away from zero, half up", New(-5, USD), 1, 2, RoundHalfUp, -3},
		{"negative halfway rounds to even quotient, half even", New(-5, USD), 1, 2, RoundHalfEven, -2},
		{"non-zero remainder always rounds up", New(100, USD), 1, 3, RoundUp, 34},
		{"non-zero remainder always truncates down", New(100, USD), 1, 3, RoundDown, 33},
		{"negative remainder rounds up away from zero", New(-100, USD), 1, 3, RoundUp, -34},
		{"negative remainder truncates toward zero", New(-100, USD), 1, 3, RoundDown, -33},
		{"zero amount stays zero", Zero(USD), 825, 10000, RoundHalfUp, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.amount.Mul(tc.num, tc.denom, tc.mode)
			if err != nil {
				t.Fatalf("Mul returned error: %v", err)
			}
			if got.Units() != tc.want {
				t.Errorf("Mul(%d, %d) = %d, want %d", tc.num, tc.denom, got.Units(), tc.want)
			}
			if got.Currency() != tc.amount.Currency() {
				t.Errorf("Mul currency = %s, want %s", got.Currency(), tc.amount.Currency())
			}
		})
	}
}

func TestMulErrors(t *testing.T) {
	if _, err := New(100, USD).Mul(1, 0, RoundHalfUp); err == nil {
		t.Fatalf("Mul with zero denominator: expected an error, got nil")
	}

	if _, err := New(100, USD).Mul(1, -4, RoundHalfUp); err == nil {
		t.Fatalf("Mul with negative denominator: expected an error, got nil")
	}

	if _, err := New(math.MaxInt64, USD).Mul(2, 1, RoundHalfUp); err == nil {
		t.Fatalf("Mul overflow: expected an error, got nil")
	}
}
