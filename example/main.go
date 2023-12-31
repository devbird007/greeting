package main

import (
	"fmt"

	"github.com/devbird007/greeting/greet"
)

func main() {
	greet.Hello()

	fmt.Println(greet.Shark)

	oct := greet.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}
	fmt.Println(oct.String())

	oct.Reset()
	fmt.Println(oct.String())
}
