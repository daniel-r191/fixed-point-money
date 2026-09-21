package money

import (
	"encoding/json"
	"fmt"
)

// amountJSON is the wire representation: minor units plus a currency
// code, rather than a decimal string. Encoding the units directly
// avoids a decimal round trip through Parse/String on every marshal.
type amountJSON struct {
	Units    int64  `json:"units"`
	Currency string `json:"currency"`
}

// MarshalJSON encodes the amount as its minor-unit count and currency
// code, e.g. {"units":1050,"currency":"USD"}.
func (a Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(amountJSON{Units: a.units, Currency: a.currency.Code})
}

// UnmarshalJSON decodes an amount previously written by MarshalJSON.
// The currency code must be one of the currencies this package knows
// about (see byCode in money.go); an unrecognized code is an error
// rather than a zero-precision guess.
func (a *Amount) UnmarshalJSON(data []byte) error {
	var raw amountJSON
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c, ok := byCode[raw.Currency]
	if !ok {
		return fmt.Errorf("money: unknown currency code %q", raw.Currency)
	}
	a.units = raw.Units
	a.currency = c
	return nil
}
