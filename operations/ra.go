package operations

import "fmt"

func RA(StackA *[]int) {
	if len(*StackA) >= 2 {
		first := (*StackA)[0]
		*StackA = (*StackA)[1:]
		*StackA = append(*StackA, first)
		fmt.Println("RA")
	}
}
