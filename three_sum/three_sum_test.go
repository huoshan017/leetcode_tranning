package three_sum

import (
	"fmt"
	"testing"
)

func threeSum(array []int, sum int) [][]int {
	if len(array) < 3 {
		return nil
	}

	twoSum2Index := make(map[int][2]int)
	var a [][]int
	for i := 0; i < len(array); i++ {
		left := sum - array[i]
		if idx, o := twoSum2Index[left]; o {
			a = append(a, []int{i, idx[0], idx[1]})
		} else {
			if i > 1 {
				for j := i - 1; j >= 0; j-- {
					twoSum := array[i] + array[j]
					twoSum2Index[twoSum] = [2]int{j, i}
				}
			}
		}
	}
	return a
}

func TestThreeSum(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20}
	sum := threeSum(array, -10)
	t.Logf("the sum is %v", sum)

	if len(sum) > 0 {
		c := 1
		for i := 0; i < len(sum); i++ {
			t.Logf("%v: ", c)
			var s string
			for j := 0; j < len(sum[i]); j++ {
				s += fmt.Sprintf("%v ", array[sum[i][j]])
			}
			t.Logf("  %v ", s)
			c += 1
		}
	}
}
