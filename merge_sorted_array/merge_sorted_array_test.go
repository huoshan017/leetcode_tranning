package merge_sorted_array

import "testing"

func TestMergeSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	mergeSortedArray(nums1, m, nums2, n)
	t.Logf("merged array: %v", nums1)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	m = 1
	n = 1
	mergeSortedArray(nums1, m, nums2, n)
	t.Logf("merged array: %v", nums1)
}
