package money

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAmountMarshalJSON(t *testing.T) {
	data, err := json.Marshal(New(1050, USD))
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}
	want := `{"units":1050,"currency":"USD"}`
	if string(data) != want {
		t.Errorf("Marshal = %s, want %s", data, want)
	}
}

func TestAmountUnmarshalJSON(t *testing.T) {
	var a Amount
	if err := json.Unmarshal([]byte(`{"units":1234,"currency":"KWD"}`), &a); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if a.Units() != 1234 {
		t.Errorf("Units() = %d, want 1234", a.Units())
	}
	if a.Currency() != KWD {
		t.Errorf("Currency() = %s, want KWD", a.Currency())
	}
}

func TestAmountUnmarshalJSONUnknownCurrency(t *testing.T) {
	var a Amount
	err := json.Unmarshal([]byte(`{"units":100,"currency":"XYZ"}`), &a)
	if err == nil {
		t.Fatal("Unmarshal with unknown currency: expected an error, got nil")
	}
	if !strings.Contains(err.Error(), "XYZ") {
		t.Errorf("error %q does not mention the unrecognized code", err)
	}
}

func TestAmountJSONRoundTrip(t *testing.T) {
	cases := []Amount{
		New(999, JPY),
		New(-500, EUR),
		New(0, BHD),
		New(1, GBP),
	}

	for _, want := range cases {
		data, err := json.Marshal(want)
		if err != nil {
			t.Fatalf("Marshal(%s) returned error: %v", want, err)
		}
		var got Amount
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal(%s) returned error: %v", data, err)
		}
		if got.Units() != want.Units() || got.Currency() != want.Currency() {
			t.Errorf("round trip of %s = %s, want %s", want, got, want)
		}
	}
}

func TestAmountMarshalJSONInStruct(t *testing.T) {
	type invoice struct {
		Total Amount `json:"total"`
	}

	in := invoice{Total: New(4999, USD)}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	var out invoice
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("Unmarshal returned error: %v", err)
	}
	if out.Total.Units() != 4999 || out.Total.Currency() != USD {
		t.Errorf("round trip via struct field = %s, want 49.99 USD", out.Total)
	}
}
