package money

import "testing"

func TestAllocate(t *testing.T) {
	cases := []struct {
		name   string
		amount Amount
		ratios []int
		want   []int64 // expected minor units per share
	}{
		{
			name:   "splits evenly into thirds with the remainder cent going to the first share",
			amount: New(100, USD),
			ratios: []int{1, 1, 1},
			want:   []int64{34, 33, 33},
		},
		{
			name:   "five cents three ways",
			amount: New(5, USD),
			ratios: []int{1, 1, 1},
			want:   []int64{2, 2, 1},
		},
		{
			name:   "weighted two-to-one split",
			amount: New(10, USD),
			ratios: []int{1, 2},
			want:   []int64{4, 6},
		},
		{
			name:   "a zero ratio gets nothing",
			amount: New(10, USD),
			ratios: []int{1, 0, 1},
			want:   []int64{5, 0, 5},
		},
		{
			name:   "negative amount follows the same remainder rule in reverse",
			amount: New(-100, USD),
			ratios: []int{1, 1, 1},
			want:   []int64{-34, -33, -33},
		},
		{
			name:   "a single ratio returns the whole amount",
			amount: New(999, JPY),
			ratios: []int{1},
			want:   []int64{999},
		},
		{
			name:   "zero amount produces zero shares",
			amount: Zero(USD),
			ratios: []int{1, 1},
			want:   []int64{0, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.amount.Allocate(tc.ratios...)
			if err != nil {
				t.Fatalf("Allocate returned error: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %d shares, want %d", len(got), len(tc.want))
			}

			var sum int64
			for i, share := range got {
				if share.units != tc.want[i] {
					t.Errorf("share %d = %d, want %d", i, share.units, tc.want[i])
				}
				if share.currency != tc.amount.currency {
					t.Errorf("share %d currency = %s, want %s", i, share.currency, tc.amount.currency)
				}
				sum += share.units
			}
			if sum != tc.amount.units {
				t.Errorf("shares sum to %d, want %d (the original amount)", sum, tc.amount.units)
			}
		})
	}
}

func TestAllocateErrors(t *testing.T) {
	cases := []struct {
		name   string
		amount Amount
		ratios []int
	}{
		{"no ratios at all", New(100, USD), nil},
		{"a negative ratio", New(100, USD), []int{1, -1}},
		{"ratios that all sum to zero", New(100, USD), []int{0, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := tc.amount.Allocate(tc.ratios...); err == nil {
				t.Fatalf("expected an error, got nil")
			}
		})
	}
}
