package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genFirst first.go
func (r *Repository) genFirst(rt reflect.Type, abbr, filename, paths, modelName string) error {
	data := struct {
		Package     string
		GenQueryPkg string
		RepoPkg     string
		ModelPkg    string
		ModelName   string
		RepoPkgName string
		StructName  string
		Abbr        string
	}{
		Package:     filename,
		GenQueryPkg: r.genQueryPkg,
		RepoPkg:     r.repoPkg,
		ModelPkg:    rt.PkgPath(),
		ModelName:   modelName,
		RepoPkgName: r.repoPkgName,
		StructName:  rt.Name(),
		Abbr:        abbr,
	}

	file, err := os.Create(path.Join(paths, "first.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genFirstTemplate()).Parse(r.genFirstTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
