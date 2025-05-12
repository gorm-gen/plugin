package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genLast last.go
func (r *Repository) genLast(rt reflect.Type, abbr, filename, paths, modelName string) error {
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

	file, err := os.Create(path.Join(paths, "last.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genLastTemplate()).Parse(r.genLastTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
