package main

import "fmt"

func one() {
	defer fmt.Println("deferred in one()")
	two()
}

func two() {
	defer fmt.Println("deferred in two()")
	three()
}

func three() {
	panic("This call stack's too deep from me!")
}

func main() {
	one()
}
