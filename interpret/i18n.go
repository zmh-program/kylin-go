package interpret

func (i *Interpreter) UseCall() interface{} {
	if token := i.Peek(); token.Type == String {
		lang := token.Value

		i.i18n.Register(i.scope, lang)
	} else {
		i.Throw("SyntaxError", "Use must have a language")
	}

	return nil
}
