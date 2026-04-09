package operations

func SA(StackA *[]int) {
	var tempStackA []int
	if len(*StackA) >= 2 {
		first := (*StackA)[0]
		second := (*StackA)[1]
		*StackA = (*StackA)[2:]
		tempStackA = append(tempStackA, second, first)
		tempStackA = append(tempStackA, *StackA...)
		*StackA = tempStackA

	}
}
