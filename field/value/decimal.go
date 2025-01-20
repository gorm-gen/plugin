package value

import (
	"database/sql/driver"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	value decimal.Decimal
}

func NewDecimal(value decimal.Decimal) *Decimal {
	return &Decimal{value: value}
}

func (d Decimal) Value() (driver.Value, error) {
	return d.value.InexactFloat64(), nil
}
