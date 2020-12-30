package main

import (
	"fmt"
)

type Liters float64
type Milliliters float64
type Gallons float64

func Example_defined_type() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)

	fmt.Println(carFuel, busFuel)

	//Output:
	//10 240
}

func Example_type_conversion() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)

	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

	//Output:
	//Gallons: 10.6 Liters: 238.5
}

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}
func Example_method1() {
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()

	//Output:
	//Hi! a MyType value
	//Hi! another value
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func Example_method_with_parameters() {
	value := MyType("MyType value")
	value.MethodWithParameters(4, true)

	//Output:
	//MyType value
	//4
	//true
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func Example_double_pointer() {
	number := Number(4)
	fmt.Println("Original value:", number)
	number.Double() //Go가 자동으로 receiver를 pointer로 변환하여 호출해줌
	fmt.Println("number after calling Double:", number)

	//Output:
	//Original value: 4
	//number after calling Double: 8
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func Example_method_pointer_receiver_parameter() {
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod() //value가 자동으로 pointer로 변환된다

	pointer.method() //pointer의 값도 자동으로 얻어온다 (*이렇게 할 필요가 없음)
	pointer.pointerMethod()
	//Output:
	//Method with value receiver
	//Method with pointer receiver
	//Method with value receiver
	//Method with pointer receiver
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func Example_method2() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	//Output:
	//2.000 liters equals 0.528 gallons
	//500.000 milliliters equals 0.132 gallons
}
