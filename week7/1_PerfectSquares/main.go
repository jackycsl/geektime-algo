package main

import "math"

// O(n * √n) time
// O(n) space
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt32
	}
	f[0] = 0
	// f[j] = min(f[j], f[j - i * i] + 1)
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = min(f[j], f[j-i*i]+1)
		}
	}
	return f[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
