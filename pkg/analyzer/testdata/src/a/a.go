package a

// Incorrect doc does something // want `doc comment should start with the function name "IncorrectDoc"`
func IncorrectDoc() {}

// CorrectDoc starts with the function name
func CorrectDoc() {}

// Incorrect Type is a type // want `doc comment should start with the type name "IncorrectType"`
type IncorrectType struct{}

// CorrectType starts with the type name
type CorrectType struct{}

// Incorrect Var is a variable // want `doc comment should start with the variable name "IncorrectVar"`
var IncorrectVar int

// CorrectVar starts with the variable name
var CorrectVar int

// Incorrect Const is a constant // want `doc comment should start with the constant name "IncorrectConst"`
const IncorrectConst = 1

// CorrectConst starts with the constant name
const CorrectConst = 1
