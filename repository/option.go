package repository

type Option func(*Repository)

func WithModule(module string) Option {
	return func(r *Repository) {
		r.module = module
	}
}

func WithRepoPath(repoPath string) Option {
	return func(r *Repository) {
		r.repoPath = repoPath
	}
}

func WithGenQueryPkg(genQueryPkg string) Option {
	return func(r *Repository) {
		r.genQueryPkg = genQueryPkg
	}
}

func WithGormDBVar(gormDBVar string) Option {
	return func(r *Repository) {
		r.gormDBVar = gormDBVar
	}
}

func WithGormDBVarPkg(gormDBVarPkg string) Option {
	return func(r *Repository) {
		r.gormDBVarPkg = gormDBVarPkg
	}
}

func WithZapVar(zapVar string) Option {
	return func(r *Repository) {
		r.zapVar = zapVar
	}
}

func WithZapVarPkg(zapVarPkg string) Option {
	return func(r *Repository) {
		r.zapVarPkg = zapVarPkg
	}
}
