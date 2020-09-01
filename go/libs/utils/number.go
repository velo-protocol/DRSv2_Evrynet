package utils

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func StringToAmount(number string) (*big.Int, error) {
	d, err := decimal.NewFromString(number)
	if err != nil {
		return nil, err
	}

	d = d.Shift(5)
	b := new(big.Int)
	b, _ = b.SetString(d.String(), 10)

	return b, nil
}

func AmountToString(amount *big.Int) string {
	d := decimal.NewFromInt(amount.Int64())
	d = d.Shift(-5).Truncate(5)
	return d.StringFixed(5)
}

func IsDecimalValid(s decimal.Decimal) bool {
	d1 := s.Shift(5)
	d2 := d1.Truncate(0)
	if !d1.Equal(d2) {
		return false
	}
	return true
}
