package nums_same_consec_diff

import (
	"math"
)

/**
// 返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。

// 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。

// 你可以按 任何顺序 返回答案。

// 示例 1：

// 输入：n = 3, k = 7
// 输出：[181,292,707,818,929]
// 解释：注意，070 不是一个有效的数字，因为它有前导零。
*/

type pair struct {
	val [2]int
}

// 求出差值为k的所有两个数字的集合
func getTwoNumWithDiff(k int) []pair {
	var pairList []pair
	// 先求出差的绝对值为k的两个数字 i,j
	for i := 0; i <= 9; i++ {
		if i+k <= 9 {
			pairList = append(pairList, pair{val: [2]int{i, i + k}})
		}
	}
	return pairList
}

// 这个函数的关键问题在于每一个非负整数只有两个数字组合而成，也就是求出仅由这两个差值为k的数字组成的n位数
func NumsSameConsecDiff(n int, k int) []int {
	if n <= 1 || k < 0 || k > 9 {
		return nil
	}

	pairList := getTwoNumWithDiff(k)

	var nums []int
	for _, pair := range pairList {
		val := [2]int{pair.val[0], pair.val[1]}
		// 从低位到高位遍历累加计算
		for i := 2; i <= n; i++ {
			if pair.val[(i+1)%2] > 0 {
				val[0] += pair.val[(i+1)%2] * int(math.Pow10(i-1))
			} else if i == n {
				val[0] = 0
			}
			if pair.val[i%2] > 0 {
				val[1] += pair.val[i%2] * int(math.Pow10(i-1))
			} else if i == n {
				val[1] = 0
			}
		}
		for i := 0; i < len(val); i++ {
			if val[i] > 0 {
				nums = append(nums, val[i])
			}
		}
	}

	return nums
}
