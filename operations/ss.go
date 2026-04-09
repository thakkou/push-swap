package operations

import "fmt"

func SS(StackA *[]int, StackB *[]int) {
	SA(StackA)
	SB(StackB)
	fmt.Println("SS")
}
