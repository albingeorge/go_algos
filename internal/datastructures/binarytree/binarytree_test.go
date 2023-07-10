package binarytree

import "testing"

func TestBinaryTreePreorder(t *testing.T) {
	tree := BinaryTree{}
	tree.Root = &TreeNode{
		Data: 1,
		Left: &TreeNode{
			Data: 2,
			Left: &TreeNode{
				Data: 6,
			},
			Right: &TreeNode{
				Data: 4,
			},
		},
		Right: &TreeNode{
			Data: 3,
			Left: &TreeNode{
				Data: 5,
			},
			Right: &TreeNode{
				Data: 7,
			},
		},
	}

	got := tree.PreorderTraversal()
	res := []int{1, 2, 6, 4, 3, 5, 7}
	for i, v := range got {
		if v != res[i] {
			t.Errorf("expected: %v, received: %v", res, got)
			return
		}
	}
}
