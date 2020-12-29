package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99 //(*s).rate 하지 않고 dot operator으로 struct에 접근하여 지정할 수 있다
}

func Example_struct_포인트_사용으로_값_변경() {
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)

	//Output:
	//4.99
}

func Example_int_pointer1() {
	var value int = 2
	var pointer *int = &value
	fmt.Println(pointer)

	//Output:
}

func Example_int_pointer2() {
	var value int = 2
	var pointer *int = &value
	fmt.Println(*pointer) //포인터의 값을 출력한다

	//Output:
	//2
}

type myStruct struct {
	myField int
}

func Example_struct_pointer1() {
	var value myStruct
	value.myField = 3
	//var pointer *myStruct = &value
	//fmt.Println(*pointer.myField) //invalid indirect of pointer.myField (type int)
	//Output:
}

func Example_struct_pointer2() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println((*pointer).myField)

	//Output:
	//3
}

func Example_struct_pointer3() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value

	//위와 같이 할 필요없이 dot operator로 접근하면 struct의 값을 * 사용없이 접근할 수 있다
	pointer.myField = 9 //struct field에 값을 assign할때도 * 사용없이 할 수 있다
	fmt.Println(pointer.myField)

	//Output:
	//9
}
