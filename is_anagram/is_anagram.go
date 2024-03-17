package is_anagram

/**
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false


提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母


进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xar9lv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func isAnagram(s, t string) bool {
	l := len(s)
	l2 := len(t)

	if l == l2 && l == 0 {
		return true
	}

	if l != l2 {
		return false
	}

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	arr1, arr2 := []byte(s), []byte(t)

	for i := 0; i < l; i++ {
		v, o := m1[arr1[i]]
		if !o {
			m1[arr1[i]] = 1
		} else {
			m1[arr1[i]] = v + 1
		}
		v, o = m2[arr2[i]]
		if !o {
			m2[arr2[i]] = 1
		} else {
			m2[arr2[i]] = v + 1
		}
	}

	for i := 0; i < l; i++ {
		a := arr1[i]
		if m1[a] != m2[a] {
			return false
		}
	}

	return true
}
