package main

import (
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/japanese"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	dansk.Hej()

	japanese.Konichiwa()
}
