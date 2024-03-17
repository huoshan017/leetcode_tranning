package max_product

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	result := maxProduct(nums)
	t.Logf("result: %v", result)
	nums = []int{-2, 0, -1}
	result = maxProduct(nums)
	t.Logf("result: %v", result)
	nums = []int{0, 2}
	result = maxProduct(nums)
	t.Logf("result: %v", result)
	nums = []int{-3, 0, 1, -2}
	result = maxProduct(nums)
	t.Logf("result: %v", result)
}
