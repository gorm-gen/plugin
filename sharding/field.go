package sharding

import (
	"gorm.io/gen/field"
)

func NewInt(genField field.Int) field.Int {
	return field.NewInt("", genField.ColumnName().String())
}

func NewInt8(genField field.Int8) field.Int8 {
	return field.NewInt8("", genField.ColumnName().String())
}

func NewInt16(genField field.Int16) field.Int16 {
	return field.NewInt16("", genField.ColumnName().String())
}

func NewInt32(genField field.Int32) field.Int32 {
	return field.NewInt32("", genField.ColumnName().String())
}

func NewInt64(genField field.Int64) field.Int64 {
	return field.NewInt64("", genField.ColumnName().String())
}

func NewUint(genField field.Uint) field.Uint {
	return field.NewUint("", genField.ColumnName().String())
}

func NewUint8(genField field.Uint8) field.Uint8 {
	return field.NewUint8("", genField.ColumnName().String())
}

func NewUint16(genField field.Uint16) field.Uint16 {
	return field.NewUint16("", genField.ColumnName().String())
}

func NewUint32(genField field.Uint32) field.Uint32 {
	return field.NewUint32("", genField.ColumnName().String())
}

func NewUint64(genField field.Uint64) field.Uint64 {
	return field.NewUint64("", genField.ColumnName().String())
}

func NewString(genField field.String) field.String {
	return field.NewString("", genField.ColumnName().String())
}

func NewField(genField field.Field) field.Field {
	return field.NewField("", genField.ColumnName().String())
}

func NewTime(genField field.Time) field.Time {
	return field.NewTime("", genField.ColumnName().String())
}

func NewFloat32(genField field.Float32) field.Float32 {
	return field.NewFloat32("", genField.ColumnName().String())
}

func NewFloat64(genField field.Float64) field.Float64 {
	return field.NewFloat64("", genField.ColumnName().String())
}

func NewBool(genField field.Bool) field.Bool {
	return field.NewBool("", genField.ColumnName().String())
}
