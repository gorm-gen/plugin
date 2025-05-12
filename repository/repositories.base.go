package repository

import (
	"html/template"
	"os"
	"path"
)

// repositoriesBase 生成repositories/base.gen.go文件
func (r *Repository) repositoriesBase() error {
	data := struct {
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

	file, err := os.Create(path.Join(r.repoPath, "base.gen.go"))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	t, err := template.New(r.baseTemplate()).Parse(r.baseTemplate())
	if err != nil {
		return err
	}
	if err = t.Execute(file, data); err != nil {
		return err
	}
	return nil
}
