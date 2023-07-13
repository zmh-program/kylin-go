package interpret

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

const (
	EOF TokenType = iota
	Integer
	Float
	String
	True
	False
	Null
	Identifier
	Addition
	Subtraction
	Multiplication
	Division
	Modulo
	Exponent
	LeftParenthesis
	RightParenthesis
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace
	Comma
	Period
	Colon
	Semicolon
	Equals
	PlusEquals
	MinusEquals
	TimesEquals
	DividedEquals
	ModuloEquals
	ExponentEquals
	And
	Or
	Not
	IsEquals
	NotEquals
)
