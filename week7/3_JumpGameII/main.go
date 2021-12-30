package main

import "math"

// 贪心
// O(n) time
// O(1) space
func jump(nums []int) int {
	now := 0
	ans := 0
	for now < len(nums)-1 {
		right := now + nums[now]
		// [now + 1, right]
		if right >= len(nums)-1 {
			return ans + 1
		}
		nextRight := right
		next := now
		for i := now + 1; i <= right; i++ {
			if i+nums[i] > nextRight {
				nextRight = i + nums[i]
				next = i
			}
		}
		now = next
		ans++
	}
	return ans
}

// 贪心会比较快，因为可能不必遍历所有数组，next right 可以跳过一些数组。

// 动规
// O(n) time
// O(n) space
func jumpdp(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = 0

	for i := 1; i < n; i++ {
		f[i] = math.MaxInt32
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
