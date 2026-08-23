package money

// Currency describes the minor-unit precision of an ISO 4217 currency.
// Minor is the number of digits after the decimal point used by that
// currency's minor unit: 2 for USD, 0 for JPY, 3 for KWD.
type Currency struct {
	Code  string
	Minor int
}

func (c Currency) String() string {
	return c.Code
}

// A handful of currencies covering the common case (2 decimal places)
// and the edge cases that break code written only against USD: zero
// decimal places (JPY) and three decimal places (KWD, BHD).
var (
	USD = Currency{Code: "USD", Minor: 2}
	EUR = Currency{Code: "EUR", Minor: 2}
	GBP = Currency{Code: "GBP", Minor: 2}
	JPY = Currency{Code: "JPY", Minor: 0}
	KWD = Currency{Code: "KWD", Minor: 3}
	BHD = Currency{Code: "BHD", Minor: 3}
)
