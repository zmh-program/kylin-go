package module

import (
	"kylin/include"
)

type Module struct {
	name    string
	path    string
	runtime KyRuntime
}

func NewModule(name string, path string, runtime KyRuntime) *Module {
	return &Module{
		name:    name,
		path:    path,
		runtime: runtime,
	}
}

func (m *Module) Run() {
	m.runtime.Run()
}

func (m *Module) GetRuntime() KyRuntime {
	return m.runtime
}

func (m *Module) GetName() string {
	return m.name
}

func (m *Module) GetPath() string {
	return m.path
}

func (m *Module) GetScope() *include.Scope {
	return m.runtime.GetScope()
}

func (m *Module) SetScope(scope *include.Scope) {
	m.runtime.SetScope(scope)
}
