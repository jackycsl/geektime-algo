package main

// O(log mn) time
// O(1) space
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	// from top to bottom, find element matrix[x][0] that <= target
	l, r := 0, m-1
	for l < r {
		mid := l + r + 1>>1
		if matrix[mid][0] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	row := r
	if matrix[row][0] == target {
		return true
	}
	if matrix[row][0] > target {
		return false
	}

	// from left to right
	l, r = 0, n-1
	for l < r {
		mid := l + r + 1>>1
		if matrix[row][mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	col := r

	return matrix[row][col] == target
}
