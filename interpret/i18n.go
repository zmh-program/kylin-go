package interpret

import "kylin/lexer"

func (i *KyRuntime) UseCall() interface{} {
	if token := i.Peek(); token.Type == lexer.String {
		lang := token.Value

		i.i18n.Register(i.scope, lang)
	} else {
		i.Throw("SyntaxError", "Use must have a language")
	}

	return nil
}
