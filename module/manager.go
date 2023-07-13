package module

type Manager struct {
	modules map[string]*Module
}

func NewManager() *Manager {
	return &Manager{
		modules: make(map[string]*Module),
	}
}

func (m *Manager) Add(module *Module) {
	m.modules[module.GetName()] = module
}

func (m *Manager) Has(name string) bool {
	if _, ok := m.modules[name]; ok {
		return true
	}
	return false
}

func (m *Manager) Import(name string, _interpret Interpreter) {
	module := NewModule(name, name, _interpret)
	module.Run()
	m.Add(module)
}

func (m *Manager) Get(name string) *Module {
	return m.modules[name]
}

func (m *Manager) Delete(name string) {
	delete(m.modules, name)
}

func (m *Manager) Clear() {
	m.modules = make(map[string]*Module)
}

func (m *Manager) Copy() *Manager {
	manager := NewManager()
	for _, v := range m.modules {
		manager.Add(v)
	}
	return manager
}

func (m *Manager) GetModules() map[string]*Module {
	return m.modules
}
