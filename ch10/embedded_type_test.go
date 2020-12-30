package main

import (
	"github.com/headfirstgo/mypackage"
)

func Example_exported_methods() {
	value := mypackage.MyType{}
	value.ExportedMethod()
	//value.exportedMethod() //접근이 안됨

	//Output:
	//Hi from ExportedMethod on EmbeddedType
}
