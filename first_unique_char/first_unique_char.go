package first_unique_char

/**
字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xaph0j/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// 用map来保存每个字符的第一个索引和出现次数
func firstUniqueChar(s string) int {
	l := len(s)
	if l == 0 {
		return -1
	}
	if l == 1 {
		return 0
	}

	m := make(map[byte]int)
	arr := []byte(s)
	for i := 0; i < l; i++ {
		a := arr[i]
		// 只有一个值保存索引（大于等于0），超过一个值保存个数
		if v, o := m[a]; !o {
			m[a] = i
		} else {
			// 已经有一个值，再加一个变成-2
			if v >= 0 {
				m[a] = -2
			} else {
				m[a] = v - 1
			}
		}
	}

	for i := 0; i < l; i++ {
		a := arr[i]
		if m[a] >= 0 {
			return m[a]
		}
	}

	return -1
}
