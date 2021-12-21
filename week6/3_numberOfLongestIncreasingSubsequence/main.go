package main

// O(n^2) time
// O(n) space
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	f := make([]int, n)
	count := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = 1
		count[i] = 1
	}

	maxCount := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if f[j]+1 > f[i] {
					f[i] = f[j] + 1
					count[i] = count[j]
				} else if f[j]+1 == f[i] { //repeat
					count[i] += count[j]
				}
			}
			if f[i] > maxCount {
				maxCount = f[i]
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		if maxCount == f[i] {
			result += count[i]
		}
	}

	return result
}
