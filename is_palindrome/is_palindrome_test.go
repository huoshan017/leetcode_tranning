package is_palindrome

import "testing"

func TestIsPalinDrome(t *testing.T) {
	var numList = []int{1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 6, 9, 6, 7, 5, 4, 4, 3, 3, 2, 2, 2, 1}
	var list *ListNode
	var node *ListNode
	for i := 0; i < len(numList); i++ {
		tmp := &ListNode{
			numList[i],
			nil,
		}
		if list == nil {
			list = tmp
		} else {
			node.Next = tmp
		}
		node = tmp
	}
	t.Logf("result: %v", isPalindrome(list))
}

func TestIsPalinDromeStr(t *testing.T) {
	str := "A man, a plan, a canal: Panama"
	result := isPalindromeStr(str)
	t.Logf("result: %v", result)
	str = "0P"
	result = isPalindromeStr(str)
	t.Logf("result: %v", result)
}
