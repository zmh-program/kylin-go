package interpret

import "fmt"

type Scope struct {
	parent    *Scope
	variables map[string]interface{}
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent:    parent,
		variables: make(map[string]interface{}),
	}
}

func (s *Scope) Get(name string) interface{} {
	if val, ok := s.variables[name]; ok {
		return val
	}
	if s.parent != nil {
		return s.parent.Get(name)
	}
	panic(fmt.Sprintf("Variable '%s' not found", name))
}

func (s *Scope) Set(name string, value interface{}) {
	s.variables[name] = value
}

func (s *Scope) Has(name string) bool {
	if _, ok := s.variables[name]; ok {
		return true
	}
	if s.parent != nil {
		return s.parent.Has(name)
	}
	return false
}

func (s *Scope) Delete(name string) {
	delete(s.variables, name)
}

func (s *Scope) Clear() {
	s.variables = make(map[string]interface{})
}

func (s *Scope) Copy() *Scope {
	scope := NewScope(nil)
	for k, v := range s.variables {
		scope.Set(k, v)
	}
	return scope
}
