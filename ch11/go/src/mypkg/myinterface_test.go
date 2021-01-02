package mypkg

import "fmt"

func Example_myinteface() {
	var value MyInterface
	value = MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())

	//Output:
	//MethodWithoutParameters called
	//MethodWithParameter called with 127.3
	//Hi from MethodWithReturnValue
}
