package termutil

import "os"

type Option func(t *Terminal)

func WithLogFile(path string) Option {
	return func(t *Terminal) {
		t.logFile, _ = os.Create(path)
	}
}

func WithTheme(theme *Theme) Option {
	return func(t *Terminal) {
		t.theme = theme
	}
}
