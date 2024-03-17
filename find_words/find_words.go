package find_words

/**
单词搜索 II
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。



示例 1：


输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
输出：["eat","oath"]
示例 2：


输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]


提示：

m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] 是一个小写英文字母
1 <= words.length <= 3 * 104
1 <= words[i].length <= 10
words[i] 由小写英文字母组成
words 中的所有字符串互不相同


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xaorig/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func findWords(board [][]byte, words []string) []string {
	m := make(map[byte][]int)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			c := board[i][j]
			idx := i*len(board[0]) + j
			if v, o := m[c]; !o {
				m[c] = []int{idx}
			} else {
				m[c] = append(v, idx)
			}
		}
	}
	var str []string
	for i := 0; i < len(words); i++ {
		w := []byte(words[i])
		for n := 0; n < len(w); n++ {
			var b []int
			var o bool
			if b, o = m[w[n]]; !o {
				return []string{}
			}
			for r := 0; r < len(b); r++ {
				h, l := idx2hl(board, r)
				if w[n] == board[h][l] {

				}
			}
		}
	}
	return str
}

func hl2idx(board [][]byte, h, l int) int {
	return h*len(board) + l
}

func idx2hl(board [][]byte, idx int) (int, int) {
	h := idx / len(board[0])
	l := idx % len(board[0])
	return h, l
}
