package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genCount count.go
func (r *Repository) genMultiList(rt reflect.Type, abbr, filename, paths, shardingKey, shardingKeyType, shardingKeyTypeFormat, modelName string) error {
	data := struct {
		Package               string
		GenQueryPkg           string
		RepoPkg               string
		RepoPkgName           string
		StructName            string
		ModelPkg              string
		ModelName             string
		Abbr                  string
		ShardingKey           string
		ShardingKeyType       string
		ShardingKeyTypeFormat string
		ChanSign              template.HTML
		DecimalPkg            template.HTML
		ToShardingValue       template.HTML
		ShardingValueTo       string
	}{
		Package:               filename,
		GenQueryPkg:           r.genQueryPkg,
		RepoPkg:               r.repoPkg,
		RepoPkgName:           r.repoPkgName,
		StructName:            rt.Name(),
		ModelPkg:              rt.PkgPath(),
		ModelName:             modelName,
		Abbr:                  abbr,
		ShardingKey:           shardingKey,
		ShardingKeyType:       shardingKeyType,
		ShardingKeyTypeFormat: shardingKeyTypeFormat,
		ChanSign:              template.HTML("<-"),
		ToShardingValue:       "k",
		ShardingValueTo:       "shardingValue := v.ShardingValue",
	}

	if shardingKeyType != "string" {
		_typeStart := shardingKeyType + "("
		_typeEnd := shardingKeyType + ")"
		if shardingKeyType == "int64" {
			_typeStart = ""
			_typeEnd = ""
		}
		data.DecimalPkg = `
    "github.com/shopspring/decimal"`
		data.ToShardingValue = `fmt.Sprintf("%d", k)`
		data.ShardingValueTo = `_shardingValue, _ := decimal.NewFromString(v.ShardingValue)
					shardingValue := ` + _typeStart + `_shardingValue.BigInt().Int64()` + _typeEnd
	}

	file, err := os.Create(path.Join(paths, "multi.list.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genMultiListTemplate()).Parse(r.genMultiListTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
