package main

func Example_panics() {
	panic("oh, no, we're going down")
	//Output:
}

func one() {
	two()
}

func two() {
	three()
}

func three() {
	panic("This call stack's too deep from me!")
}

func Example_one() {
	one()

	//Output:
}
