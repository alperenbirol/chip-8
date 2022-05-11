package stack

import "errors"

var errStackEmpty = errors.New("stack is empty")

type stack []uint16

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Pop() (uint16, error) {
	if len(*s) == 0 {
		return 0, errStackEmpty
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *stack) Push(v uint16) {
	*s = append(*s, v)
}
