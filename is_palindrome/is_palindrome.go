package is_palindrome

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	_, half := endHalfOfList(head)
	rlist := revertList(half)

	for rlist != nil && head != nil {
		if rlist.Val != head.Val {
			return false
		}
		rlist = rlist.Next
		head = head.Next
	}

	revertList(head)

	return true
}

// 反转链表
func revertList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func endHalfOfList(head *ListNode) (end, half *ListNode) {
	end = head
	half = head
	for end.Next != nil && half.Next != nil {
		end = end.Next
		half = half.Next
		if end.Next != nil {
			end = end.Next
		}
	}
	return
}

// 是否回文字符串
func isPalindromeStr(s string) bool {
	l := len(s)
	a := []byte(s)
	// 定义两个索引指向开头和最后
	i, j := 0, l-1
	for i <= j {
		if !isdigit(a[i]) && !isletter(a[i]) {
			i += 1
			continue
		}
		if !isdigit(a[j]) && !isletter(a[j]) {
			j -= 1
			continue
		}
		if i > j {
			break
		}
		ai := a[i]
		if isupper(a[i]) {
			ai, _ = tolower(a[i])
		}
		aj := a[j]
		if isupper(a[j]) {
			aj, _ = tolower(a[j])
		}
		if ai != aj {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}

func isdigit(d byte) bool {
	return int(d) >= '0' && int(d) <= '9'
}

func isupper(c byte) bool {
	return int(c) >= 'A' && int(c) <= 'Z'
}

func islower(c byte) bool {
	return int(c) >= 'a' && int(c) <= 'z'
}

func isletter(c byte) bool {
	return islower(c) || isupper(c)
}

func tolower(c byte) (byte, bool) {
	if !isupper(c) {
		return 0, false
	}
	return byte(int(c) + (int('a') - int('A'))), true
}
