package main

// O(n) time
// O(n) space
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2

	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}
