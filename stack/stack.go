package stack

import (
	"strconv"
	"strings"
)

type Stack struct {
	Items []int
}

func (stack *Stack) Parse(s string) (*Stack, error) {
	data := strings.Split(s, " ")
	stackA := Stack{}
	for _, d := range data {
		if d == "" {
			continue
		}
		n, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		stackA.Push(n)
	}
	return &stackA, nil
}

func (stack *Stack) Push(n int) {
	stack.Items = append(stack.Items, n)
}
