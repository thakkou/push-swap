package operations

func PA(StackA *[]int, StackB *[]int) {
	if len(*StackB) > 0 {
		top := (*StackB)[0]
		*StackB = (*StackB)[1:]
		*StackA = append(*StackA, top)
	}
}
