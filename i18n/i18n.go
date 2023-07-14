package i18n

type I18n struct {
	Language string
	Keyword  map[string]string
	Builtin  map[string]string
}

type Manager struct {
	Sources map[string]I18n
}
