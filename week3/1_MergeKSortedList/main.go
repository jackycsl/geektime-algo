package main

import "sort"

// O(n×n!) Time
// O(n) Space
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool, len(nums))
	ans := [][]int{}
	n := len(nums)
	tmp := []int{}

	var recur func(int)
	recur = func(pos int) {
		if pos == n {
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		for i, v := range nums {
			if used[i] || i > 0 && !used[i-1] && v == nums[i-1] {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = true
			recur(pos + 1)
			used[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}

	recur(0)
	return ans
}
