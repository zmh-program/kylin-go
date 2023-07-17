package lexer

const (
	IdentifierSequence = iota
	WhileSequence
	ForSequence
	IfSequence
	TrySequence
	AssignSequence
	FunctionSequence
	FunctionCallSequence
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
	Variable string
	Array    ExecSequence
	Body     []ExecSequence
}

type ElifStruct struct { // elif condition {}
	Condition []Token
	Body      []ExecSequence
}

type IfStruct struct { // if condition {} <ElifStruct> else {}
	Condition []Token
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
	Value []Value
}
