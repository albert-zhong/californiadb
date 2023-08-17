package db

type Operator interface {
	Open()
	Next()
	Close()
	Schema()
}

type Filter struct {
	// =
	// >
	// <
}

type BooleanPredicate interface {
	Evaluate(other bool)
}

type StringPredicate interface {
	Evaluate(other bool)
}

type IntegerPredicate interface {
	Evaluate(other int32)
}

type BooleanEqualsPredicate struct {
}

type StringEqualsPredicate struct {
}

type IntegerEqualsPredicate struct {
}

type IntegerGreaterPredicate struct {
}

type IntegerLesserPredicate struct {
}

type Limit struct {
}

type Join struct {
}
