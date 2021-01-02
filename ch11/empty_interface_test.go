package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle) //use a type assertion to get a Whistle
	if ok {
		whistle.MakeSound()
	}
}

func Example_empty_interface1() {
	AcceptAnything(3.1415)
	AcceptAnything(Whistle("Toyco Canary"))

	//Output:
	//3.1415
	//Toyco Canary
	//Tweet!
}
