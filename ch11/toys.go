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

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	play(Robot("Botco Ambler"))
}
