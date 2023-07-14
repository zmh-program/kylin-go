package i18n

import (
	"fmt"
	"kylin/i18n/source"
	"kylin/include"
)

type I18n struct {
	Language string
	Keyword  map[string]string
	Builtin  map[string]string
}

type Manager struct {
	Sources map[string]I18n
	Lang    string
}

func NewManager() *Manager {
	return &Manager{
		Sources: map[string]I18n{
			"chinese": {
				Language: "chinese",
				Keyword:  source.ChineseKeyword,
				Builtin:  source.ChineseBuiltin,
			},
		},
		Lang: "",
	}
}

func (m *Manager) Set(lang string) bool {
	if _, ok := m.Sources[lang]; ok {
		m.Lang = lang
		return true
	}
	return false
}

func (m *Manager) Get() string {
	return m.Lang
}

func (m *Manager) Has(lang string) bool {
	_, ok := m.Sources[lang]
	return ok
}

func (m *Manager) HasKeyword(keyword string) bool {
	if m.Lang == "" {
		return false
	}
	_, ok := m.Sources[m.Lang].Keyword[keyword]
	return ok
}

func (m *Manager) GetKeyword(key string) string {
	if m.Lang == "" {
		return key
	}
	return m.Sources[m.Lang].Keyword[key]
}

func (m *Manager) GetBuiltin(key string) string {
	if m.Lang == "" {
		return key
	}
	return m.Sources[m.Lang].Builtin[key]
}

func (m *Manager) Register(scope *include.Scope, language string) {
	scope.Set("i18n", map[string]interface{}{
		"set": func(lang string) bool {
			return m.Set(lang)
		},
		"get": func() string {
			return m.Get()
		},
	})

	for key, value := range m.Sources[language].Builtin {
		if fn, ok := scope.Get(value); ok {
			scope.Set(key, fn)
		}

		fmt.Println(key, value)
	}
}
