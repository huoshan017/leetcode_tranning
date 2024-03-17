package remove_duplicates

// 原地删除有序数组中的重复元素，并返回删除后的数组的新长度
func remove_duplicates(array []int32) int32 {
	var length = int32(len(array))
	var c int32
	if length > 0 {
		c = 1
	}
	var fi, si int32
	for ; fi < length; fi++ {
		if array[si] == array[fi] {
			continue
		}
		si += 1
		array[si] = array[fi]
		c += 1
	}
	return c
}
