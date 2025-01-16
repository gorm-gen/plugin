package generate

import (
	_ "github.com/shopspring/decimal"
	_ "gorm.io/plugin/soft_delete"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type Generate struct {
	generate *gen.Generator
}

func New(generate *gen.Generator) *Generate {
	return &Generate{generate: generate}
}

func (g *Generate) SetDataTypeMap(newMap map[string]func(columnType gorm.ColumnType) (dataType string)) {
	dataMap := map[string]func(detailType gorm.ColumnType) string{
		"decimal": func(detailType gorm.ColumnType) (dataType string) {
			return "decimal.Decimal"
		},
		"datetime": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}
			return "*time.Time"
		},
		"timestamp": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}
			return "*time.Time"
		},
		"tinyint": func(detailType gorm.ColumnType) (dataType string) {
			return "int8"
		},
		"smallint": func(detailType gorm.ColumnType) (dataType string) {
			return "int16"
		},
		"mediumint": func(detailType gorm.ColumnType) (dataType string) {
			return "int32"
		},
		"int": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "soft_delete.DeletedAt"
			}
			return "int"
		},
		"bigint": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "soft_delete.DeletedAt"
			}
			return "int64"
		},
		"varchar": func(detailType gorm.ColumnType) (dataType string) {
			if nullable, ok := detailType.Nullable(); nullable && ok {
				return "*string"
			}
			return "string"
		},
		"char": func(detailType gorm.ColumnType) (dataType string) {
			if nullable, ok := detailType.Nullable(); nullable && ok {
				return "*string"
			}
			return "string"
		},
	}

	if newMap != nil {
		for k, v := range newMap {
			if v == nil {
				delete(dataMap, k)
				continue
			}
			dataMap[k] = v
		}
	}

	g.generate.WithDataTypeMap(dataMap)
}

func (g *Generate) SetJSONTagNameStrategy(newTags map[string]string) {
	tags := map[string]string{
		"created_at": "sql_datetime",
		"updated_at": "sql_datetime",
	}
	if newTags != nil {
		for k, v := range newTags {
			if v == "" {
				delete(tags, k)
				continue
			}
			tags[k] = v
		}
	}
	g.generate.WithJSONTagNameStrategy(func(columnName string) string {
		if v, ok := tags[columnName]; ok {
			return columnName + "\" time_format:\"" + v
		}
		return columnName
	})
}
