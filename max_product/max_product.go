package max_product

func maxProduct(nums []int) int {
	var max int
	f := false
	for i := 0; i < len(nums); i++ {
		if !f {
			max = nums[i]
			f = true
		}
		product := nums[i]
		if product == 0 {
			continue
		}
		if product > max {
			max = product
		}
		for j := i + 1; j < len(nums); j++ {
			if product*nums[j] > max {
				max = product * nums[j]
			}
			product *= nums[j]
			if product == 0 {
				break
			}
		}
	}
	return max
}
