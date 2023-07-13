package module

import (
	"kylin/include"
)

type Module struct {
	name        string
	path        string
	interpreter Interpreter
}

func NewModule(name string, path string, _interpret Interpreter) *Module {
	return &Module{
		name:        name,
		path:        path,
		interpreter: _interpret,
	}
}

func (m *Module) Run() {
	m.interpreter.Run()
}

func (m *Module) GetInterpreter() Interpreter {
	return m.interpreter
}

func (m *Module) GetName() string {
	return m.name
}

func (m *Module) GetPath() string {
	return m.path
}

func (m *Module) GetScope() *include.Scope {
	return m.interpreter.GetScope()
}

func (m *Module) SetScope(scope *include.Scope) {
	m.interpreter.SetScope(scope)
}
