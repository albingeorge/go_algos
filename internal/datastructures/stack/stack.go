package stack

import "fmt"

type StackNode struct {
	Next *StackNode
	Data int
}

type Stack struct {
	Root *StackNode
	Size int
}

func (s *Stack) Insert(data int) {
	n := StackNode{
		Data: data,
	}

	if s.Root == nil {
		s.Root = &n
	} else {
		currNode := s.Root
		for currNode.Next != nil {
			currNode = currNode.Next
		}

		currNode.Next = &n
	}

	s.Size++
}

func (s *Stack) Pop() int {
	var val int
	if s.Root == nil {
		panic("out of range")
	} else if s.Size == 1 {
		val = s.Root.Data
		s.Root = nil
	} else {
		currNode := s.Root
		for currNode.Next.Next != nil {
			currNode = currNode.Next
		}
		val = currNode.Next.Data
		currNode.Next = nil
	}
	s.Size--
	return val
}

func (s *Stack) Print() {
	if s.Root == nil {
		fmt.Println("empty")
		return
	}

	curr := s.Root
	for curr != nil {
		fmt.Printf("%v ", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}
