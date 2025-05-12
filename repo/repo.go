package repo

import (
	"html/template"
	"os"
	"path"
	"reflect"
	"strings"
)

type Repo struct {
	module       string
	repoPath     string
	repoPkg      string
	repoPkgName  string
	genQueryPkg  string
	gormDBVar    string
	gormDBVarPkg string
	zapVar       string
	zapVarPkg    string
}

func New(opts ...Option) *Repo {
	repo := &Repo{
		module:       "demo",
		repoPath:     "internal/repositories",
		repoPkgName:  "repositories",
		genQueryPkg:  "demo/internal/query",
		gormDBVar:    "global.DB",
		gormDBVarPkg: "demo/internal/global",
		zapVar:       "global.Logger",
		zapVarPkg:    "demo/internal/global",
	}

	for _, opt := range opts {
		opt(repo)
	}

	repoPathArr := strings.Split(repo.repoPath, "/")
	repo.repoPkgName = repoPathArr[len(repoPathArr)-1]

	repo.repoPkg = path.Join(repo.module, repo.repoPath)

	return repo
}

func (r *Repo) Generate(models ...interface{}) error {
	if len(models) == 0 {
		return nil
	}

	if err := os.MkdirAll(r.repoPath, os.ModePerm); err != nil {
		return err
	}

	baseData := struct {
		Package      string
		GormDBVarPkg string
		GenQueryPkg  string
		GormDBVar    string
	}{
		Package:      r.repoPkgName,
		GormDBVarPkg: r.gormDBVarPkg,
		GenQueryPkg:  r.genQueryPkg,
		GormDBVar:    r.gormDBVar,
	}

	// 1、生成repo.base文件
	baseFile, err := os.Create(path.Join(r.repoPath, "base.gen.go"))
	if err != nil {
		return err
	}
	defer baseFile.Close()
	t, err := template.New(r.baseTemplate()).Parse(r.baseTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(baseFile, baseData); err != nil {
		return err
	}

	for _, model := range models {
		rt := reflect.TypeOf(model)
		if rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}

		abbr := strings.ToLower(rt.Name()[:1])

		filename := abbr + rt.Name()[1:]
		paths := path.Join(r.repoPath, filename)
		if err = os.MkdirAll(paths, os.ModePerm); err != nil {
			return err
		}

		modelPkgArr := strings.Split(rt.PkgPath(), "/")
		modelName := modelPkgArr[len(modelPkgArr)-1]

		// base.go
		{
			genBaseData := struct {
				Package     string
				ZapVarPkg   string
				GenQueryPkg string
				RepoPkg     string
				RepoPkgName string
				StructName  string
				Abbr        string
			}{
				Package:     filename,
				ZapVarPkg:   r.zapVarPkg,
				GenQueryPkg: r.genQueryPkg,
				RepoPkg:     r.repoPkg,
				RepoPkgName: r.repoPkgName,
				StructName:  rt.Name(),
				Abbr:        abbr,
			}

			baseFile, err = os.Create(path.Join(paths, "base.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genBaseTemplate()).Parse(r.genBaseTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genBaseData); err != nil {
				return err
			}
		}

		// count.go
		{
			genCountData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "count.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genCountTemplate()).Parse(r.genCountTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genCountData); err != nil {
				return err
			}
		}

		// create.go
		{
			genCreateData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "create.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genCreateTemplate()).Parse(r.genCreateTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genCreateData); err != nil {
				return err
			}
		}

		// delete.go
		{
			genDeleteData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "delete.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genDeleteTemplate()).Parse(r.genDeleteTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genDeleteData); err != nil {
				return err
			}
		}

		// first.go
		{
			genFirstData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "first.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genFirstTemplate()).Parse(r.genFirstTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genFirstData); err != nil {
				return err
			}
		}

		// last.go
		{
			genLastData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "last.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genLastTemplate()).Parse(r.genLastTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genLastData); err != nil {
				return err
			}
		}

		// list.go
		{
			genListData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "list.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genListTemplate()).Parse(r.genListTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genListData); err != nil {
				return err
			}
		}

		// take.go
		{
			genTakeData := struct {
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

			baseFile, err = os.Create(path.Join(paths, "take.gen.go"))
			if err != nil {
				return err
			}
			defer baseFile.Close()
			t, err = template.New(r.genTakeTemplate()).Parse(r.genTakeTemplate())
			if err != nil {
				return err
			}
			if err = t.Execute(baseFile, genTakeData); err != nil {
				return err
			}
		}
	}

	return nil
}
