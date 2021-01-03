package main

import (
	"fmt"
	"log"
)

func Socialize() {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}

func Example_defer() {
	Socialize()

	//Output:
	//Hello!
	//Nice weather, eh?
	//Goodbye
}

func SocializeWithError() error {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}

func Example_defer2() {
	err := SocializeWithError()
	if err != nil {
		log.Fatal(err)
	}

	//Output:
}
