package main

import "math"

// O(n^2) time
// O(n^2) space

// [2]
// [3, 4]
// [6, 5, 7]
// [4, 1, 8, 3]
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		// first element of each row, the value can only be the previous row first element
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		// last element of each row, can only be the last element of previous row
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[n-1][i])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
