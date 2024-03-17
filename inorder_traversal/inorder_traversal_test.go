package inorder_traversal

import "testing"

func TestInorderTraversal(t *testing.T) {
	nodeList := []*TreeNode{{Val: 1}, nil, {Val: 2}, {Val: 3}}
	root := nodeList[0]
	root.Left = nodeList[1]
	root.Right = nodeList[2]
	root.Right.Left = nodeList[3]
	inorderTraversal(root, func(node *TreeNode) {
		if node == nil {
			t.Logf(" ")
		} else {
			t.Logf("%v ", node.Val)
		}
	})
}
