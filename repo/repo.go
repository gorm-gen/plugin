package repo

import (
	"html/template"
	"os"
	"path"
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
}

func New(opts ...Option) *Repo {
	repo := &Repo{
		module:       "demo",
		repoPath:     "internal/repositories",
		repoPkgName:  "repositories",
		genQueryPkg:  "demo/internal/query",
		gormDBVar:    "global.DB",
		gormDBVarPkg: "demo/internal/global",
	}

	for _, opt := range opts {
		opt(repo)
	}

	repoPathArr := strings.Split(repo.repoPath, "/")
	repo.repoPkgName = repoPathArr[len(repoPathArr)-1]

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

	return nil
}
