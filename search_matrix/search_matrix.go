package search_matrix

func _search(matrix [][]int, left, right, top, bottom int, target int) bool {
	if top > bottom || left > right {
		return false
	}

	mv := (top + bottom) / 2
	mh := (left + right) / 2

	result := false
	if matrix[mv][mh] == target {
		result = true
	} else if matrix[mv][mh] < target {
		// 先找后面再找下面
		if !_search(matrix, mh+1, right, top, bottom, target) {
			result = _search(matrix, left, right, mv+1, bottom, target)
		} else {
			result = true
		}
	} else {
		// 先找前面再找上面
		if !_search(matrix, left, mh-1, top, bottom, target) {
			result = _search(matrix, left, right, top, mv-1, target)
		} else {
			result = true
		}
	}

	return result
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	left, right := 0, n-1
	top, bottom := 0, m-1
	return _search(matrix, left, right, top, bottom, target)
}

func searchSortedMatrixSimple(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	// 从左下角或右上角开始查找，类似于二分查找
	h := 0
	v := m - 1
	for v >= 0 && h <= n-1 {
		if matrix[v][h] == target {
			return true
		} else if matrix[v][h] < target {
			// 先往右移慢慢接近
			h += 1
		} else {
			// 再往上
			v -= 1
		}
	}
	return false
}
