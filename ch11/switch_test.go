package main

import "fmt"

type Switch string

type Toggleable interface {
	toggle()
}

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

func Example_switch() {
	s := Switch("off")

	//var t Toggleable = s
	var t Toggleable = &s
	t.toggle()
	t.toggle()

	//Output:
	//on
	//off
}
