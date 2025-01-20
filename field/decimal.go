package field

import (
	"github.com/shopspring/decimal"
	"gorm.io/gen/field"
)

type Field struct {
	tableName string
	column    string
}

func NewDecimal(genField field.Field, opts ...Option) *Field {
	f := &Field{
		column: genField.ColumnName().String(),
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

type Option func(*Field)

func WithTableName(tableName string) Option {
	return func(f *Field) {
		f.tableName = tableName
	}
}

func (f *Field) newField64() field.Float64 {
	return field.NewFloat64(f.tableName, f.column)
}

func (f *Field) toFloat64(val decimal.Decimal) float64 {
	return val.InexactFloat64()
}

// Value =
func (f *Field) Value(val decimal.Decimal) field.AssignExpr {
	return f.newField64().Value(f.toFloat64(val))
}

// Sum 累加
func (f *Field) Sum() field.Float64 {
	return f.newField64().Sum()
}

// Add +=
func (f *Field) Add(val decimal.Decimal) field.AssignExpr {
	return f.newField64().Add(f.toFloat64(val))
}

// Sub -=
func (f *Field) Sub(val decimal.Decimal) field.AssignExpr {
	return f.newField64().Sub(f.toFloat64(val))
}

// Mul *=
func (f *Field) Mul(val decimal.Decimal) field.AssignExpr {
	return f.newField64().Mul(f.toFloat64(val))
}

// Div /=
func (f *Field) Div(val decimal.Decimal) field.AssignExpr {
	return f.newField64().Div(f.toFloat64(val))
}
