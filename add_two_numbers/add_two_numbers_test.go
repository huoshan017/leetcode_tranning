package add_two_numbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	nodeList1 := []*ListNode{
		{2, nil}, {4, nil}, {3, nil}, {8, nil},
	}
	nodeList2 := []*ListNode{
		{5, nil}, {6, nil}, {4, nil}, {9, nil},
	}

	for i := 0; i < len(nodeList1); i++ {
		if l1 == nil {
			l1 = nodeList1[i]
		} else {
			nodeList1[i-1].Next = nodeList1[i]
		}
	}

	for i := 0; i < len(nodeList2); i++ {
		if l2 == nil {
			l2 = nodeList2[i]
		} else {
			nodeList2[i-1].Next = nodeList2[i]
		}
	}

	t.Logf("l1: ")
	ll1 := l1
	for ll1 != nil {
		t.Logf(" %v", *ll1)
		ll1 = ll1.Next
	}

	t.Logf("l2: ")
	ll2 := l2
	for ll2 != nil {
		t.Logf(" %v", *ll2)
		ll2 = ll2.Next
	}

	l := addTwoNumbers(l1, l2)

	t.Logf("result: ")

	for l != nil {
		t.Logf(" %v", *l)
		l = l.Next
	}
}
