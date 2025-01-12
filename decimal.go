package bybit

import "github.com/shopspring/decimal"

type NullDecimalV2 decimal.NullDecimal

//goland:noinspection GoMixedReceiverTypes
func (d *NullDecimalV2) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" || string(decimalBytes) == `""` {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.UnmarshalJSON(decimalBytes)
}

//goland:noinspection GoMixedReceiverTypes
func (d NullDecimalV2) N() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}
