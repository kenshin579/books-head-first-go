package main

import "fmt"

type Whistle string
type Horn string
type Robot string

type NoiseMaker interface {
	MakeSound()
}

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func play(n NoiseMaker) {
	n.MakeSound()
	// Walk is not part of the NoiseMaker interface.
	// Uncommenting the next line will cause a compile error!
	// n.Walk()
}

func Example_noisemaker() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Robot("Botco Ambler"))

	//Output:
	//Tweet!
	//Honk!
	//Beep Boop
}

func Example_walk() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot) //convert back to the concrete type using a type assertion
	robot.Walk()

	//Output:
	//Beep Boop
	//Powering legs

}
