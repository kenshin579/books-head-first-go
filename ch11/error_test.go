package main

import (
	"fmt"
	"log"
)

type OverheatError float64 //float64이지만, error interface를 만족시킨다

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func Example_error1() {
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}

	//Output:
}

//Go builtin erorr interface
//type error interface {
//	Error() string
//}

type ComedyError string

func (c ComedyError) Error() string { //error interface를 만족함
	return string(c) //do a type conversion
}

func Example_error2() {
	var err error
	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	//Output:
	//What's a programmer's favorite beer? Logger!
}
