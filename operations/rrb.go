package operations

func RRB(StackB *[]int) {
	var tempStackB []int
	if len(*StackB) >= 2 {
		Last := (*StackB)[len(*StackB)-1]
		*StackB = (*StackB)[:len(*StackB)-1]
		tempStackB = append(tempStackB, Last)
		tempStackB = append(tempStackB, *StackB...)
		*StackB = tempStackB

	}
}
