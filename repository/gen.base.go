package repository

import (
	"html/template"
	"os"
	"path"
	"reflect"
)

type Condition string

type Update string

type Order string

func (r *Repository) genBase(rt reflect.Type, abbr, filename, paths string) error {
	data := struct {
		Package     string
		ZapVarPkg   string
		GenQueryPkg string
		RepoPkg     string
		RepoPkgName string
		StructName  string
		Abbr        string
		Conditions  []Condition
		Updates     []Update
		Orders      []Order
	}{
		Package:     filename,
		ZapVarPkg:   r.zapVarPkg,
		GenQueryPkg: r.genQueryPkg,
		RepoPkg:     r.repoPkg,
		RepoPkgName: r.repoPkgName,
		StructName:  rt.Name(),
		Abbr:        abbr,
		Conditions:  r.genConditionOpt(rt, abbr),
	}

	file, err := os.Create(path.Join(paths, "base.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.genBaseTemplate()).Parse(r.genBaseTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
