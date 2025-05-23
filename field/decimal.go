package field

import (
	"github.com/shopspring/decimal"
	"gorm.io/gen/field"
)

type Field struct {
	tableName string
	column    string
}

func NewDecimal(genField field.Expr, opts ...Option) *Field {
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

// Lte <=
func (f *Field) Lte(val decimal.Decimal) field.Expr {
	return f.newField64().Lte(f.toFloat64(val))
}

// Lt <
func (f *Field) Lt(val decimal.Decimal) field.Expr {
	return f.newField64().Lt(f.toFloat64(val))
}

// Gte >=
func (f *Field) Gte(val decimal.Decimal) field.Expr {
	return f.newField64().Gte(f.toFloat64(val))
}

// Gt >
func (f *Field) Gt(val decimal.Decimal) field.Expr {
	return f.newField64().Gt(f.toFloat64(val))
}

// Eq =
func (f *Field) Eq(val decimal.Decimal) field.Expr {
	return f.newField64().Eq(f.toFloat64(val))
}

// Neq !=
func (f *Field) Neq(val decimal.Decimal) field.Expr {
	return f.newField64().Neq(f.toFloat64(val))
}

// In in
func (f *Field) In(vals ...decimal.Decimal) field.Expr {
	var float64Vals []float64
	for _, val := range vals {
		float64Vals = append(float64Vals, f.toFloat64(val))
	}
	return f.newField64().In(float64Vals...)
}

// NotIn not in
func (f *Field) NotIn(vals ...decimal.Decimal) field.Expr {
	var float64Vals []float64
	for _, val := range vals {
		float64Vals = append(float64Vals, f.toFloat64(val))
	}
	return f.newField64().NotIn(float64Vals...)
}

// Between between
func (f *Field) Between(left, right decimal.Decimal) field.Expr {
	return f.newField64().Between(f.toFloat64(left), f.toFloat64(right))
}

// NotBetween not between
func (f *Field) NotBetween(left, right decimal.Decimal) field.Expr {
	return f.newField64().NotBetween(f.toFloat64(left), f.toFloat64(right))
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
