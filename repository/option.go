package repository

type Option func(*Repository)

func WithModule(module string) Option {
	return func(r *Repository) {
		r.module = module
	}
}

// WithRepositoryPath 生成仓库所在路径
func WithRepositoryPath(repositoryPath string) Option {
	return func(r *Repository) {
		r.repoPath = repositoryPath
	}
}

// WithGenQueryPkg gen的query.Query所在的包路径名
func WithGenQueryPkg(genQueryPkg string) Option {
	return func(r *Repository) {
		r.genQueryPkg = genQueryPkg
	}
}

// WithGormDBVar 本地存储的Gorm.DB的变量
func WithGormDBVar(gormDBVar string) Option {
	return func(r *Repository) {
		r.gormDBVar = gormDBVar
	}
}

// WithGormDBVarPkg 本地存储的Gorm.DB的变量所在的包路径名
func WithGormDBVarPkg(gormDBVarPkg string) Option {
	return func(r *Repository) {
		r.gormDBVarPkg = gormDBVarPkg
	}
}

// WithZapVar 本地存储的zap的变量
func WithZapVar(zapVar string) Option {
	return func(r *Repository) {
		r.zapVar = zapVar
	}
}

// WithZapVarPkg 本地存储的zap的变量所在的包路径名
func WithZapVarPkg(zapVarPkg string) Option {
	return func(r *Repository) {
		r.zapVarPkg = zapVarPkg
	}
}
