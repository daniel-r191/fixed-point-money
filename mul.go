package money

import (
	"errors"
	"fmt"
	"math/big"
)

// RoundingMode controls how Mul resolves a fractional result down to a
// whole number of minor units.
type RoundingMode int

const (
	// RoundHalfUp rounds a result exactly halfway between two minor
	// units away from zero. This is the rounding most tax authorities
	// expect.
	RoundHalfUp RoundingMode = iota
	// RoundHalfEven rounds a result exactly halfway between two minor
	// units to whichever of the two is even (banker's rounding), which
	// avoids the slight upward bias RoundHalfUp accumulates over many
	// transactions.
	RoundHalfEven
	// RoundUp always rounds a non-zero remainder away from zero.
	RoundUp
	// RoundDown always truncates a non-zero remainder toward zero.
	RoundDown
)

// Mul multiplies the amount by the rational factor num/denom, rounding
// the result to a whole number of minor units according to mode. The
// factor is expressed as a fraction rather than a float so that rates
// like an 8.25% tax (num=825, denom=10000) or a 15% discount taken off
// (num=85, denom=100) never introduce the binary rounding error a
// float64 multiplier would.
func (a Amount) Mul(num, denom int64, mode RoundingMode) (Amount, error) {
	if denom <= 0 {
		return Amount{}, errors.New("money: mul denominator must be positive")
	}

	product := new(big.Int).Mul(big.NewInt(a.units), big.NewInt(num))
	bigDenom := big.NewInt(denom)

	q, r := new(big.Int), new(big.Int)
	q.QuoRem(product, bigDenom, r)
	q = roundQuotient(q, r, bigDenom, mode)

	if !q.IsInt64() {
		return Amount{}, fmt.Errorf("money: overflow multiplying %s by %d/%d", a, num, denom)
	}
	return Amount{units: q.Int64(), currency: a.currency}, nil
}

// roundQuotient adjusts the truncated quotient q of product/denom,
// given the truncated remainder r, according to mode. denom is assumed
// positive, which makes r's sign match product's sign (Go's big.Int
// division truncates toward zero).
func roundQuotient(q, r, denom *big.Int, mode RoundingMode) *big.Int {
	if r.Sign() == 0 || mode == RoundDown {
		return q
	}

	neg := r.Sign() < 0
	if mode == RoundUp {
		return bumpAwayFromZero(q, neg)
	}

	twiceR := new(big.Int).Lsh(new(big.Int).Abs(r), 1)
	switch twiceR.Cmp(denom) {
	case -1:
		return q
	case 1:
		return bumpAwayFromZero(q, neg)
	default: // exactly halfway
		if mode == RoundHalfUp {
			return bumpAwayFromZero(q, neg)
		}
		if new(big.Int).Abs(q).Bit(0) == 1 { // q is odd
			return bumpAwayFromZero(q, neg)
		}
		return q
	}
}

func bumpAwayFromZero(q *big.Int, neg bool) *big.Int {
	if neg {
		return new(big.Int).Sub(q, big.NewInt(1))
	}
	return new(big.Int).Add(q, big.NewInt(1))
}
