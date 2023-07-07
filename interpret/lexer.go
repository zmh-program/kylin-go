package interpret

type TokenType int

const (
	Addition TokenType = iota
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
	NotEquals
	Assignment
	PlusEquals
	MinusEquals
	TimesEquals
	DividedEquals
	ModuloEquals
	ExponentEquals
	And
	Or
	Not
)
