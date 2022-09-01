package stack

import "errors"

var errStackEmpty = errors.New("stack is empty")

type Stack []uint16

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (uint16, error) {
	if len(*s) == 0 {
		return 0, errStackEmpty
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *Stack) Push(v uint16) {
	*s = append(*s, v)
}

func (s *Stack) Peek() (uint16, error) {
	if len(*s) == 0 {
		return 0, errStackEmpty
	}
	return (*s)[len(*s)-1], nil
}
