package merge_sorted_array

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// 特殊情况，m=0
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	// 先从最后一个开始比较
	i1 := m - 1
	i2 := n - 1
	i := m + n - 1
	for i1 >= 0 && i2 >= 0 {
		// 相等时把数组2的值放在后面
		if nums1[i1] <= nums2[i2] {
			nums1[i] = nums2[i2]
			// 数组2的索引往前移
			i2 -= 1
		} else {
			nums1[i] = nums1[i1]
			i1 -= 1
		}
		i -= 1
	}

	// 表示数组nums1都被排好，nums2还有一部分数据没有放到nums1
	if i1 < 0 {
		for ; i>=0; i-- {
			nums1[i] = nums2[i2]
			i2 -= 1
		}
	}
}
