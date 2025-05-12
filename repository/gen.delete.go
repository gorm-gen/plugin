package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

// genDelete delete.go
func (r *Repository) genDelete(rt reflect.Type, abbr, filename, paths string) error {
	data := struct {
		Package     string
		GenQueryPkg string
		RepoPkg     string
		RepoPkgName string
		StructName  string
		Abbr        string
	}{
		Package:     filename,
		GenQueryPkg: r.genQueryPkg,
		RepoPkg:     r.repoPkg,
		RepoPkgName: r.repoPkgName,
		StructName:  rt.Name(),
		Abbr:        abbr,
	}

	file, err := os.Create(path.Join(paths, "delete.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genDeleteTemplate()).Parse(r.genDeleteTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
