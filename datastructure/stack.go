package datastructure

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack(v *list.List) *Stack {
	return &Stack{v}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}

	return nil
}

func StackDrive() {
	stack := NewStack(list.New())
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v", val)
		val = stack.Pop()
	}
}
