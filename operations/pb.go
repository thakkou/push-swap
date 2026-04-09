package operations

import "fmt"

func PB(StackA *[]int, StackB *[]int) {
	if len(*StackA) > 0 {
		top := (*StackA)[0]
		*StackA = (*StackA)[1:]
		*StackB = append([]int{top}, *StackB...)
		fmt.Println("PB")
	}
}
