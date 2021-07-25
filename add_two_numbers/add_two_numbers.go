package add_two_numbers

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	var l *ListNode
	var currNode *ListNode
	carry := false
	for p1 != nil || p2 != nil {
		var v int
		if p1 != nil && p2 != nil {
			v = p1.Val + p2.Val
		} else if p1 != nil {
			v = p1.Val
		} else if p2 != nil {
			v = p2.Val
		}
		if carry {
			v += 1
			carry = false
		}
		if v >= 10 {
			v %= 10
			carry = true
		}

		node := &ListNode{
			Val:  v,
			Next: nil,
		}

		if l == nil {
			l = node
		} else {
			currNode.Next = node
		}
		currNode = node

		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if carry {
		currNode.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return l
}
