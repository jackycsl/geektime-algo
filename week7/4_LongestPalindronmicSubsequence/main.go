package main

// O(n^2) time
// O(n^2) space
func longestPalindromeSubseq(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	// from bottom to top
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		// from left to right
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
			}
		}
	}
	return f[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
