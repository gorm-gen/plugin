package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genCount count.go
func (r *Repository) genMultiLast(rt reflect.Type, abbr, filename, paths, shardingKey, shardingKeyType, shardingKeyTypeFormat, modelName string) error {
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
	}

	file, err := os.Create(path.Join(paths, "multi.last.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genMultiLastTemplate()).Parse(r.genMultiLastTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
