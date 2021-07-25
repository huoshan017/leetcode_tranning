package revert_list

import "testing"

func outputList(head *ListNode, t *testing.T) {
	t.Logf("List: ")
	node := head
	for node != nil {
		t.Logf("%+v", node)
		node = node.next
	}
}

func getList() *ListNode {
	node8 := &ListNode{9, nil}
	node7 := &ListNode{8, node8}
	node6 := &ListNode{7, node7}
	node5 := &ListNode{6, node6}
	node4 := &ListNode{5, node5}
	node3 := &ListNode{4, node4}
	node2 := &ListNode{3, node3}
	node1 := &ListNode{2, node2}
	head := &ListNode{1, node1}
	return head
}

func TestRevertList(t *testing.T) {
	head := getList()

	t.Logf("before revert")
	outputList(head, t)
	newHead := revertList(head)
	t.Logf("after revert")
	outputList(newHead, t)
}

func revertListRecursive(head *ListNode, prev *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.next
	head.next = prev
	prev = head
	head = next
	// 只有一个节点时
	if head == nil {
		return prev
	}
	return revertListRecursive(head, prev)
}

func TestRevertListRecursive(t *testing.T) {
	head := getList()

	t.Logf("before revert recursive")
	outputList(head, t)
	newHead := revertListRecursive(head, nil)
	t.Logf("after revert")
	outputList(newHead, t)
}
