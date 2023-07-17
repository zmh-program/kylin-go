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

var TokenMap = map[TokenType]string{
	EOF:                "EOF",
	Sep:                "Sep",
	Integer:            "Integer",
	Float:              "Float",
	String:             "String",
	True:               "True",
	False:              "False",
	Null:               "Null",
	Identifier:         "Identifier",
	Array:              "Array",
	Map:                "Map",
	Function:           "Function",
	Return:             "Return",
	If:                 "If",
	Elif:               "Elif",
	Else:               "Else",
	For:                "For",
	In:                 "In",
	While:              "While",
	Break:              "Break",
	Continue:           "Continue",
	Import:             "Import",
	Use:                "Use",
	Try:                "Try",
	Catch:              "Catch",
	Addition:           "+",
	Subtraction:        "-",
	Multiplication:     "*",
	Division:           "/",
	Modulo:             "%",
	Exponent:           "^",
	LeftParenthesis:    "(",
	RightParenthesis:   ")",
	LeftBracket:        "[",
	RightBracket:       "]",
	LeftBrace:          "{",
	RightBrace:         "}",
	Comma:              ",",
	Period:             ".",
	Colon:              ":",
	Equals:             "=",
	PlusEquals:         "+=",
	MinusEquals:        "-=",
	TimesEquals:        "*=",
	DividedEquals:      "/=",
	ModuloEquals:       "%=",
	ExponentEquals:     "^=",
	And:                "&&",
	Or:                 "||",
	Not:                "!",
	IsEquals:           "==",
	NotEquals:          "!=",
	GreaterThan:        ">",
	LessThan:           "<",
	GreaterThanOrEqual: ">=",
	LessThanOrEqual:    "<=",
}

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
