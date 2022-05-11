package stack

import "errors"

var errStackEmpty = errors.New("stack is empty")

type stack []byte

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Pop() (byte, error) {
	if len(*s) == 0 {
		return 0, errStackEmpty
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

func (s *stack) Push(v byte) {
	*s = append(*s, v)
}
