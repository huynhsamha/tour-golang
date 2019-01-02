package A

// Foo : exported function, should be comment by go-lint
func Foo() {

}

func bar() {
	// this is private in package A, but other file in package A can use bar()
}

// Animal : type animal is exported, shoul be comment by go-lint
type Animal struct {
	a string // that is private
	B int    // that is exported, that mean public
}
