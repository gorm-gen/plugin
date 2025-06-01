package repository

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
)

type Repository struct {
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

func New(opts ...Option) *Repository {
	repo := &Repository{
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

func (r *Repository) Generate(models ...interface{}) error {
	if len(models) == 0 {
		return nil
	}

	if err := os.MkdirAll(r.repoPath, os.ModePerm); err != nil {
		return err
	}

	// repositories/base.go
	if err := r.repositoriesBase(); err != nil {
		return err
	}

	for _, model := range models {
		if err := r.generate(model, ""); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) ShardingGenerate(shardingStructName string, models ...interface{}) error {
	for _, model := range models {
		if err := r.generate(model, shardingStructName); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) generate(model interface{}, shardingStructName string) error {
	rt := reflect.TypeOf(model)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	abbr := strings.ToLower(rt.Name()[:1])
	filename := abbr + rt.Name()[1:]
	paths := path.Join(r.repoPath, filename)
	if err := os.MkdirAll(paths, os.ModePerm); err != nil {
		return err
	}
	modelPkgArr := strings.Split(rt.PkgPath(), "/")
	modelName := modelPkgArr[len(modelPkgArr)-1]

	// base.go
	if err := r.genBase(rt, abbr, filename, paths); err != nil {
		return err
	}

	// count.go
	if err := r.genCount(rt, abbr, filename, paths); err != nil {
		return err
	}

	// create.go
	if err := r.genCreate(rt, abbr, filename, paths, modelName); err != nil {
		return err
	}

	// delete.go
	if err := r.genDelete(rt, abbr, filename, paths); err != nil {
		return err
	}

	// first.go
	if err := r.genFirst(rt, abbr, filename, paths, modelName); err != nil {
		return err
	}

	// last.go
	if err := r.genLast(rt, abbr, filename, paths, modelName); err != nil {
		return err
	}

	// list.go
	if err := r.genList(rt, abbr, filename, paths, modelName); err != nil {
		return err
	}

	// sum.go
	if err := r.genSum(rt, abbr, filename, paths); err != nil {
		return err
	}

	// take.go
	if err := r.genTake(rt, abbr, filename, paths, modelName); err != nil {
		return err
	}

	// update.go
	if err := r.genUpdate(rt, abbr, filename, paths); err != nil {
		return err
	}

	if shardingStructName == "" {
		return nil
	}

	shardingKeyExist := false
	var shardingKeyType string
	var shardingKeyTypeFormat string

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Name != shardingStructName {
			continue
		}
		allowType := false
		typ := field.Type.String()
		if r.isInt(typ) {
			allowType = true
			shardingKeyTypeFormat = "d"
		}
		if r.isString(typ) {
			allowType = true
			shardingKeyTypeFormat = "s"
		}
		if !allowType {
			return fmt.Errorf("%s sharding key %s type %s not support", rt.Name(), shardingStructName, typ)
		}
		shardingKeyType = strings.Trim(typ, "*")
		shardingKeyExist = true
		break
	}

	if !shardingKeyExist {
		return fmt.Errorf("%s not exist sharding key %s", rt.Name(), shardingStructName)
	}

	// multi.count.go
	if err := r.genMultiCount(rt, abbr, filename, paths, shardingStructName, shardingKeyType, shardingKeyTypeFormat); err != nil {
		return err
	}

	// multi.delete.go
	if err := r.genMultiDelete(rt, abbr, filename, paths, shardingStructName, shardingKeyType, shardingKeyTypeFormat); err != nil {
		return err
	}

	// multi.first.go
	if err := r.genMultiFirst(rt, abbr, filename, paths, shardingStructName, shardingKeyType, shardingKeyTypeFormat, modelName); err != nil {
		return err
	}

	// multi.last.go
	if err := r.genMultiLast(rt, abbr, filename, paths, shardingStructName, shardingKeyType, shardingKeyTypeFormat, modelName); err != nil {
		return err
	}

	// multi.sum.go
	if err := r.genMultiSum(rt, abbr, filename, paths, shardingStructName, shardingKeyType, shardingKeyTypeFormat); err != nil {
		return err
	}

	return nil
}
