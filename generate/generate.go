package generate

import (
	"fmt"

	_ "github.com/shopspring/decimal"
	_ "gorm.io/plugin/soft_delete"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type Generate struct {
	db            *gorm.DB
	outPath       string
	mode          gen.GenerateMode
	modelPkgPath  string
	dataTypeMap   map[string]func(gorm.ColumnType) string
	jsonTagName   map[string]map[string]string
	generateModel []string
	applyBasic    []interface{}
	generator     *gen.Generator
}

type Option func(*Generate)

func WithOutPath(outPath string) Option {
	return func(g *Generate) {
		g.outPath = outPath
	}
}

func WithMode(mode gen.GenerateMode) Option {
	return func(g *Generate) {
		g.mode = mode
	}
}

func WithModelPkgPath(modelPkgPath string) Option {
	return func(g *Generate) {
		g.modelPkgPath = modelPkgPath
	}
}

func (g *Generate) SetGenerateModel(tableNames ...string) {
	_tableNames := make([]string, 0, len(tableNames))
	for _, tableName := range tableNames {
		if tableName == "" {
			continue
		}
		_tableNames = append(_tableNames, tableName)
	}
	g.generateModel = append(g.generateModel, _tableNames...)
}

func (g *Generate) SetApplyBasic(models ...interface{}) {
	_models := make([]interface{}, 0, len(models))
	for _, model := range models {
		if model == nil {
			continue
		}
		_models = append(_models, model)
	}
	g.applyBasic = append(g.applyBasic, _models...)
}

func WithDataTypeMap(dataTypeMap map[string]func(gorm.ColumnType) string) Option {
	return func(g *Generate) {
		g.dataTypeMap = dataTypeMap
	}
}

func WithReplaceDataTypeMap(dataTypeMap map[string]func(gorm.ColumnType) string) Option {
	return func(g *Generate) {
		if dataTypeMap == nil {
			g.dataTypeMap = nil
			return
		}
		if g.dataTypeMap == nil {
			g.dataTypeMap = make(map[string]func(gorm.ColumnType) string)
		}
		for k, v := range dataTypeMap {
			if v == nil {
				delete(g.dataTypeMap, k)
				continue
			}
			g.dataTypeMap[k] = v
		}
	}
}

func WithJsonTagName(jsonTagName map[string]map[string]string) Option {
	return func(g *Generate) {
		g.jsonTagName = jsonTagName
	}
}

func WithReplaceJsonTagName(jsonTagName map[string]map[string]string) Option {
	return func(g *Generate) {
		if jsonTagName == nil {
			g.jsonTagName = nil
			return
		}
		if g.jsonTagName == nil {
			g.jsonTagName = make(map[string]map[string]string)
		}
		for k, v := range jsonTagName {
			if v == nil {
				delete(g.jsonTagName, k)
				continue
			}
			g.jsonTagName[k] = v
		}
	}
}

func New(db *gorm.DB, opts ...Option) *Generate {
	g := &Generate{
		db:            db,
		outPath:       "./internal/query",
		mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		modelPkgPath:  "models",
		dataTypeMap:   dataTypeMap(),
		jsonTagName:   jsonTagName(),
		generateModel: make([]string, 0),
		applyBasic:    make([]interface{}, 0),
	}

	for _, opt := range opts {
		opt(g)
	}

	g.generator = gen.NewGenerator(gen.Config{
		OutPath:      g.outPath,
		Mode:         g.mode,
		ModelPkgPath: g.modelPkgPath,
	})

	g.generator.UseDB(db)

	return g
}

func (g *Generate) Generator() *gen.Generator {
	return g.generator
}

func (g *Generate) Execute() {
	if g.dataTypeMap != nil && len(g.dataTypeMap) > 0 {
		g.generator.WithDataTypeMap(g.dataTypeMap)
	}

	if g.jsonTagName != nil && len(g.jsonTagName) > 0 {
		g.generator.WithJSONTagNameStrategy(func(columnName string) string {
			if tag, ok := g.jsonTagName[columnName]; ok {
				for k, v := range tag {
					return fmt.Sprintf(`%s" %s:"%s`, columnName, k, v)
				}
			}
			return columnName
		})
	}

	for _, tableName := range g.generateModel {
		g.generator.GenerateModel(tableName)
	}

	if len(g.applyBasic) > 0 {
		g.generator.ApplyBasic(g.applyBasic...)
	}

	g.generator.Execute()
}

func dataTypeMap() map[string]func(gorm.ColumnType) string {
	return map[string]func(detailType gorm.ColumnType) string{
		"decimal": func(detailType gorm.ColumnType) (dataType string) {
			return "decimal.Decimal"
		},
		"datetime": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}
			if nullable, ok := detailType.Nullable(); nullable && ok {
				return "*time.Time"
			}
			return "time.Time"
		},
		"timestamp": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}
			if nullable, ok := detailType.Nullable(); nullable && ok {
				return "*time.Time"
			}
			return "time.Time"
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
}

func jsonTagName() map[string]map[string]string {
	return map[string]map[string]string{
		"created_at": {"time_format": "sql_datetime"},
		"updated_at": {"time_format": "sql_datetime"},
	}
}
