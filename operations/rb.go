package operations

import "fmt"

func RB(StackB *[]int) {
	if len(*StackB) >= 2 {
		first := (*StackB)[0]
		*StackB = (*StackB)[1:]
		*StackB = append(*StackB, first)
		fmt.Println("RB")
	}
}
