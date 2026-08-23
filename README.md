# fixed-point-money

A Go library for representing currency amounts as integers instead of
floats, and for the pieces of arithmetic on those amounts that are easy
to get wrong: adding two amounts that turn out to be in different
currencies, formatting a currency that has zero decimal places (yen)
or three (dinar), and splitting an amount into shares without losing
or inventing a cent along the way.

## Why not just use float64

`0.10 + 0.20` in `float64` is `0.30000000000000004`. That's fine for a
lot of things and it is not fine for money. This library stores every
amount as an `int64` count of the currency's minor unit (cents for
USD, nothing for JPY, thousandths for KWD), so arithmetic on it is
exact.

## Usage

```go
package main

import (
	"fmt"

	money "github.com/daniel-r191/fixed-point-money"
)

func main() {
	price, err := money.Parse("19.99", money.USD)
	if err != nil {
		panic(err)
	}

	tax, err := money.Parse("1.65", money.USD)
	if err != nil {
		panic(err)
	}

	total, err := price.Add(tax)
	if err != nil {
		panic(err)
	}

	fmt.Println(total) // 21.64 USD
}
```

### Splitting a bill three ways

The obvious way to split $10.00 three ways is $3.33, $3.33, $3.33,
which adds up to $9.99, not $10.00. `Allocate` hands out the leftover
cent instead of dropping it:

```go
bill := money.New(1000, money.USD) // $10.00
shares, err := bill.Allocate(1, 1, 1)
if err != nil {
	panic(err)
}
for _, s := range shares {
	fmt.Println(s)
}
// 3.34 USD
// 3.33 USD
// 3.33 USD
```

`Allocate` also takes weighted ratios, e.g. `Allocate(2, 1)` to split
two-to-one, and works for any currency, including ones like JPY that
have no minor unit at all.

## Currencies

Currency precision differs: USD, EUR, and GBP use 2 decimal places,
JPY uses 0, and KWD and BHD use 3. `Parse` and `String` both respect
this, and `Parse` rejects input with more precision than the currency
supports rather than silently rounding it away:

```go
_, err := money.Parse("100.5", money.JPY) // error: JPY has 0 decimal places
```

## Status

Early. The core `Amount` type, parsing, formatting, `Add`/`Sub`/`Cmp`,
and `Allocate` are here and covered by tests. Multiplication with
explicit rounding modes, JSON encoding, and a broader currency table
are not yet.

## License

MIT
