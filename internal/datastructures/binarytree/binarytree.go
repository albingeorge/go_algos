package binarytree

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func CreateNode(data int) *TreeNode {
	return &TreeNode{
		Data: data,
	}
}

func (t *BinaryTree) PreorderTraversal() []int {
	res := []int{}
	if t.Root == nil {
		fmt.Println("empty stack")
		return res
	}

	stack := treeStack{}

	stack.Insert(t.Root)

	for stack.Size != 0 {
		curr := stack.Pop()
		if curr == nil {
			continue
		}
		res = append(res, curr.Data)
		fmt.Printf("%v ", curr.Data)

		stack.Insert(curr.Right)
		stack.Insert(curr.Left)
	}

	fmt.Println()
	return res
}
