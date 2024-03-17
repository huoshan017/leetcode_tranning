package remove_duplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var array = []int32{0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 6, 7, 7, 8, 9, 9, 9, 9, 9, 9, 10, 10, 11, 12, 13, 14, 15, 15, 16, 17, 18, 19, 20, 20, 20, 20, 20}
	var l = remove_duplicates(array)
	t.Logf("removed array is %v", array[:l])

	var array2 = []int32{}
	l = remove_duplicates(array2)
	t.Logf("removed empty array is %v", array[:l])
}
