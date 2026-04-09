package operations

import "fmt"

func RR(StackA *[]int, StackB *[]int) {
	RA(StackA)
	RB(StackB)
	fmt.Println("RR")
}
