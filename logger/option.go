package logger

type Option func(*Logger)

func WithPath(path string) Option {
	return func(l *Logger) {
		l.path = path
	}
}
