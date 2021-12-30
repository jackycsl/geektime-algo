package main

// O(n) time
// O(n) space
func canJump(nums []int) bool {
	n := len(nums)
	f := make([]bool, n)
	f[0] = true

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if f[j] && j+nums[j] >= i {
				f[i] = true
				break
			}
		}
	}
	return f[n-1]
}
