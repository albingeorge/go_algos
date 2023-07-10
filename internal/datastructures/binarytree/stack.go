package binarytree

type treeStackNode struct {
	Next *treeStackNode
	Data *TreeNode
}

type treeStack struct {
	Root *treeStackNode
	Size int
}

func (s *treeStack) Insert(data *TreeNode) {
	n := treeStackNode{
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

func (s *treeStack) Pop() *TreeNode {
	var val *TreeNode
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
