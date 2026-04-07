package main

import (
	"fmt"
	"os"

	"push-swap/stack"
)

func main() {
	if len(os.Args) != 2 {
		// not sure when args > 1
		return
	}
	var a *stack.Stack
	// stackB := Stack{}
	a, err := a.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(a.Items)
}
