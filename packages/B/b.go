package B

import "github.com/huynhsamha/tour-golang/packages/A"

func todo() {
	A.Foo() // OK

	// A.bar()
	// cannot refer to unexported name A.bar

	// x := A.Animal{"123", 4}
	// implicit assignment of unexported field 'a' in A.Animal literal

	x := A.Animal{B: 123}
	// x.a = "123"
	// x.a undefined (cannot refer to unexported field or method a)

	x.B = 123
}
