package interpret

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
	return nil
}
