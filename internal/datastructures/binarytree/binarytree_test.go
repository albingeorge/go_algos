package binarytree

import "testing"

func TestBinaryTreePreorder(t *testing.T) {
	tree := BinaryTree{}
	tree.Root = &TreeNode{
		Data: 4,
		Left: &TreeNode{
			Data: 3,
		},
		Right: &TreeNode{
			Data: 5,
		},
	}

	got := tree.PreorderTraversal()
	res := []int{4, 3, 5}
	for i, v := range got {
		if v != res[i] {
			t.Errorf("expected: %v, received: %v", res, got)
			return
		}
	}
}
