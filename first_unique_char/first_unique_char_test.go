package first_unique_char

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	str := "leetcode"
	result := firstUniqueChar(str)
	t.Logf("result: %v", result)
	str = "loveleetcode"
	result = firstUniqueChar(str)
	t.Logf("result: %v", result)
}
