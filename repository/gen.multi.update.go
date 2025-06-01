package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genCount count.go
func (r *Repository) genMultiUpdate(rt reflect.Type, abbr, filename, paths, shardingKey, shardingKeyType, shardingKeyTypeFormat string) error {
	data := struct {
		Package               string
		GenQueryPkg           string
		RepoPkg               string
		RepoPkgName           string
		StructName            string
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
		Abbr:                  abbr,
		ShardingKey:           shardingKey,
		ShardingKeyType:       shardingKeyType,
		ShardingKeyTypeFormat: shardingKeyTypeFormat,
		ChanSign:              template.HTML("<-"),
	}

	file, err := os.Create(path.Join(paths, "multi.update.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genMultiUpdateTemplate()).Parse(r.genMultiUpdateTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
