package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/operations"
)

type Stack struct {
	Elements []int
}

func main() {
	input := os.Args[1]

	var StackA []int
	var StackB []int

	for _, r := range strings.Fields(input) {
		element, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Error")
			return
		}
		StackA = append(StackA, element)
	}

	operations.PB(&StackA, &StackB)
	operations.PB(&StackA, &StackB)
	operations.PB(&StackA, &StackB)

	operations.RRR(&StackA, &StackB)

	fmt.Println(StackA)
	fmt.Println(StackB)
}
