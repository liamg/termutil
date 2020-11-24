package termutil

type themeFactory struct {
	theme *Theme
}

func NewThemeFactory() *themeFactory {
	theme := DefaultTheme
	return &themeFactory{
		theme: &theme,
	}
}

func (t *themeFactory) Build() *Theme {
	return t.theme
}
