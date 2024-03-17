package inorder_traversal

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	values []*TreeNode
	top    int
}

func newStack(s int) *stack {
	return &stack{
		values: make([]*TreeNode, s),
		top:    -1,
	}
}

func (s *stack) push(node *TreeNode) bool {
	if s.top >= len(s.values)-1 {
		return false
	}
	s.top += 1
	s.values[s.top] = node
	return true
}

func (s *stack) pop() bool {
	if s.top == -1 {
		return false
	}
	s.top -= 1
	return true
}

func (s *stack) gettop() *TreeNode {
	if s.top < 0 {
		return nil
	}
	return s.values[s.top]
}

func (s *stack) clear() {
	s.top = -1
}

// 非递归中序遍历
func inorderTraversal(root *TreeNode, f func(node *TreeNode)) []int {
	s := newStack(32)
	node := root
	for node != nil {
		// 持续访问左节点
		for node.Left != nil {
			s.push(node)
			node = node.Left
		}
		node = s.gettop()
		if node == nil {
			continue
		}
		s.pop()
		f(node)
		node = node.Right
		s.push(node)
	}
	s.clear()
	return nil
}
