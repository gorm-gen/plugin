package repository

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"reflect"
)

type Condition string

type Update string

type Order string

func (r *Repository) genBase(rt reflect.Type, abbr, filename, paths string) error {
	var timePkg, decimalPkg, numberDecimalPkg bool

	condition, _timePkg, _decimalPkg, _numberDecimalPkg := r.genConditionOpt(rt, abbr)
	if _timePkg {
		timePkg = true
	}
	if _decimalPkg {
		decimalPkg = true
	}
	if _numberDecimalPkg {
		numberDecimalPkg = true
	}

	var imports []template.HTML
	if timePkg {
		imports = append(imports, `    "time"

`)
	}

	if numberDecimalPkg {
		imports = append(imports, `    f "github.com/gorm-gen/plugin/field"
    "github.com/gorm-gen/plugin/field/value"
`)
	}
	if decimalPkg {
		imports = append(imports, `    "github.com/shopspring/decimal"
`)
	}

	imports = append(imports, template.HTML(fmt.Sprintf(`    "go.uber.org/zap"
    "gorm.io/gen"
    "gorm.io/gen/field"

    "%s"

    "%s"

    "%s"`, r.zapVarPkg, r.genQueryPkg, r.repoPkg)))

	data := struct {
		Package     string
		ZapVarPkg   string
		GenQueryPkg string
		RepoPkg     string
		RepoPkgName string
		StructName  string
		Abbr        string
		Imports     []template.HTML
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
		Imports:     imports,
		Conditions:  condition,
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
