package operations

import "fmt"

func RRA(StackA *[]int) {
	var tempStackA []int
	if len(*StackA) >= 2 {
		Last := (*StackA)[len(*StackA)-1]
		*StackA = (*StackA)[:len(*StackA)-1]
		tempStackA = append(tempStackA, Last)
		tempStackA = append(tempStackA, *StackA...)
		*StackA = tempStackA
		fmt.Println("RRA")

	}
}
