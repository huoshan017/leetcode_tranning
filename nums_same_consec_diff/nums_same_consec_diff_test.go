package nums_same_consec_diff

import "testing"

func TestGetTwoNumsWithDiff(t *testing.T) {
	for i := 0; i < 9; i++ {
		pairList := getTwoNumWithDiff(i)
		t.Logf("Diff is %v, two nums: %+v", i, pairList)
	}
}

func TestNumsSameConsecDiff(t *testing.T) {
	n := 10
	k := 3
	nums := NumsSameConsecDiff(n, k)
	t.Logf("same diff %v with %v digits, nums: %v", k, n, nums)
}
