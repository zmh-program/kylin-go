package lexer

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

const (
	EOF TokenType = iota
	Sep
	Integer
	Float
	String
	True
	False
	Null
	Identifier
	Array
	Map
	Function
	Return
	If
	Elif
	Else
	For
	In
	While
	Break
	Continue
	Import
	Use
	Try
	Catch
	Addition
	Subtraction
	Multiplication
	Division
	Modulo           // %
	Exponent         // ^
	LeftParenthesis  // (
	RightParenthesis // )
	LeftBracket      // [
	RightBracket     // ]
	LeftBrace        // {
	RightBrace       // }
	Comma            // ,
	Period           // .
	Colon            // :
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
	GreaterThan
	LessThan
	GreaterThanOrEqual
	LessThanOrEqual
)

var KeywordMap = map[string]TokenType{
	"true":     True,
	"false":    False,
	"null":     Null,
	"fn":       Function,
	"return":   Return,
	"if":       If,
	"elif":     Elif,
	"else":     Else,
	"for":      For,
	"in":       In,
	"while":    While,
	"break":    Break,
	"continue": Continue,
	"import":   Import,
	"use":      Use,
	"try":      Try,
	"catch":    Catch,
}

var AssignType = []TokenType{
	Equals,
	PlusEquals,
	MinusEquals,
	TimesEquals,
	DividedEquals,
	ModuloEquals,
	ExponentEquals,
}
