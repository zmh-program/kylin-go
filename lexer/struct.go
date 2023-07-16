package lexer

const (
	IdentifierSequence = iota
	WhileSequence
	ForSequence
	IfSequence
	TrySequence
	AssignSequence
	FunctionSequence
	CountSequence
)

type ExecSequence struct {
	Type int
	Data interface{} // &struct{}
}

type WhileStruct struct { // while condition {}
	Condition []Token
	Body      []ExecSequence
}

type ForStruct struct { // for condition {}
	Condition []Token
	Body      []ExecSequence
}

type ElifStruct struct { // elif condition {}
	Condition string
	Body      []ExecSequence
}

type IfStruct struct { // if condition {} <ElifStruct> else {}
	Condition string
	Body      []ExecSequence
	Elif      []ElifStruct
	Else      []ExecSequence
}

type TryStruct struct { // try {} catch {}
	Body  []ExecSequence
	Catch []ExecSequence
}

type AssignStruct struct { // variable = value
	Variable string
	Type     TokenType
	Value    ExecSequence
}

type FunctionStruct struct { // fn name(args) {}
	Name string
	Args []string
	Body []ExecSequence
}

type FunctionCallStruct struct { // f(args)
	Name string
	Args []ExecSequence
}

type CountStruct struct { // variable + variable
	Value []Token
}
