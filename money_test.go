package money

import "testing"

// TestByCodeConsistency guards the hand-written byCode table against the
// kind of copy-paste mistake it's prone to: a map key that doesn't match
// the Currency's own Code, or a Minor value outside what any real
// currency uses.
func TestByCodeConsistency(t *testing.T) {
	for key, c := range byCode {
		if c.Code != key {
			t.Errorf("byCode[%q] holds a Currency with Code %q", key, c.Code)
		}
		if c.Minor != 0 && c.Minor != 2 && c.Minor != 3 {
			t.Errorf("byCode[%q] has Minor %d, want 0, 2, or 3", key, c.Minor)
		}
	}
}

func TestByCodeZeroDecimalCurrencies(t *testing.T) {
	for _, c := range []Currency{JPY, BIF, CLP, DJF, GNF, ISK, KMF, KRW, PYG, RWF, UGX, VND, VUV, XAF, XOF, XPF} {
		if c.Minor != 0 {
			t.Errorf("%s: Minor = %d, want 0", c.Code, c.Minor)
		}
		got, ok := byCode[c.Code]
		if !ok {
			t.Errorf("%s missing from byCode", c.Code)
			continue
		}
		if got != c {
			t.Errorf("byCode[%q] = %+v, want %+v", c.Code, got, c)
		}
	}
}

func TestByCodeThreeDecimalCurrencies(t *testing.T) {
	for _, c := range []Currency{KWD, BHD, IQD, JOD, LYD, OMR, TND} {
		if c.Minor != 3 {
			t.Errorf("%s: Minor = %d, want 3", c.Code, c.Minor)
		}
		got, ok := byCode[c.Code]
		if !ok {
			t.Errorf("%s missing from byCode", c.Code)
			continue
		}
		if got != c {
			t.Errorf("byCode[%q] = %+v, want %+v", c.Code, got, c)
		}
	}
}
