package lexer

type Value struct {
	Type  TokenType
	Value interface{}
}

type FloatType float64
type IntegerType int64
type StringType string
type BooleanType bool
type NullType struct{}
type ArrayType []Value
type MapType map[string]Value
type FunctionType struct {
	Name string
	Args []string
	Body []ExecSequence
}
