package revert_list

type ListNode struct {
	value int
	next  *ListNode
}

func revertList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	var prev *ListNode
	node := head
	for node != nil {
		next := node.next
		node.next = prev
		prev = node
		node = next
	}
	return prev
}