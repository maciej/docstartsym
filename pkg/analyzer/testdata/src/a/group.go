package a

var (
	// GoodGroupedVar1 starts with the variable name
	GoodGroupedVar1 int

	// bad grouped var comment // want `doc comment should start with the variable name "BadGroupedVar"`
	BadGroupedVar int
)

type (
	// GoodGroupedType starts with the type name
	GoodGroupedType struct{}

	// bad grouped type comment // want `doc comment should start with the type name "BadGroupedType", optionally preceded by 'A' or 'An'`
	BadGroupedType struct{}
)
