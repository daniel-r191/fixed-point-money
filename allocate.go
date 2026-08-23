package money

import "errors"

// Allocate splits an amount into shares proportional to ratios, in a
// way that never loses or invents a minor unit: the shares always sum
// back to the original amount. Any unit left over after the proportional
// split (unavoidable whenever the amount doesn't divide evenly) is
// handed out one at a time, starting from the first ratio, to whichever
// ratios are non-zero.
func (a Amount) Allocate(ratios ...int) ([]Amount, error) {
	if len(ratios) == 0 {
		return nil, errors.New("money: allocate requires at least one ratio")
	}

	total := 0
	for _, r := range ratios {
		if r < 0 {
			return nil, errors.New("money: allocate ratios must not be negative")
		}
		total += r
	}
	if total == 0 {
		return nil, errors.New("money: allocate ratios must sum to more than zero")
	}

	shares := make([]Amount, len(ratios))
	remainder := a.units
	for i, r := range ratios {
		share := a.units * int64(r) / int64(total)
		shares[i] = Amount{units: share, currency: a.currency}
		remainder -= share
	}

	step := int64(1)
	if remainder < 0 {
		step = -1
	}
	for i := 0; remainder != 0; i = (i + 1) % len(ratios) {
		if ratios[i] == 0 {
			continue
		}
		shares[i].units += step
		remainder -= step
	}

	return shares, nil
}
