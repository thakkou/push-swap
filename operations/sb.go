package operations

func SB(StackB *[]int) {
	var tempStackB []int
	if len(*StackB) >= 2 {
		first := (*StackB)[0]
		second := (*StackB)[1]
		*StackB = (*StackB)[2:]
		tempStackB = append(tempStackB, second, first)
		tempStackB = append(tempStackB, *StackB...)
		*StackB = tempStackB

	}
}
