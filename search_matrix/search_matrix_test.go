package search_matrix

import "testing"

var m1 [][]int = [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
var m2 [][]int = [][]int{{5}, {6}}

func TestSearchMatrix(t *testing.T) {
	result := searchMatrix(m1, 30)
	t.Logf("result for m1: %v", result)
	result = searchMatrix(m2, 6)
	t.Logf("result for m2: %v", result)
}

func TestSearchMatrixSimple(t *testing.T) {
	result := searchSortedMatrixSimple(m1, -30)
	t.Logf("result for m1 use simple: %v", result)
	result = searchSortedMatrixSimple(m2, -6)
	t.Logf("result for m2 use simpe: %v", result)
}
