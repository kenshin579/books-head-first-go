package main

import "fmt"

func calmDown() {
	p := recover()       //retuns an interface{} value
	err, ok := p.(error) //Error 타입으로 변환해야 함
	if ok {
		fmt.Println(err.Error())
	}
}

func freakOut() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	//panic("oh no") //when we recover, freakout returns at this point
	panic(err)
	fmt.Println("I won't be run!") //this code after the panic will never be run!
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
