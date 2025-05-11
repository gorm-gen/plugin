package repo

type Option func(*Repo)

func WithModule(module string) Option {
	return func(r *Repo) {
		r.module = module
	}
}

func WithRepoPath(repoPath string) Option {
	return func(r *Repo) {
		r.repoPath = repoPath
	}
}

func WithGenQueryPkg(genQueryPkg string) Option {
	return func(r *Repo) {
		r.genQueryPkg = genQueryPkg
	}
}

func WithGormDBVar(gormDBVar string) Option {
	return func(r *Repo) {
		r.gormDBVar = gormDBVar
	}
}

func WithGormDBVarPkg(gormDBVarPkg string) Option {
	return func(r *Repo) {
		r.gormDBVarPkg = gormDBVarPkg
	}
}

func WithZapVar(zapVar string) Option {
	return func(r *Repo) {
		r.zapVar = zapVar
	}
}

func WithZapVarPkg(zapVarPkg string) Option {
	return func(r *Repo) {
		r.zapVarPkg = zapVarPkg
	}
}
