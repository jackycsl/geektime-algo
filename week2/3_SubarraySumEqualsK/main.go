package main

// O(n) time
// O(n) space
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	// map = {preSum: count}
	mp := make(map[int]int)
	mp[0] = 1
	for _, num := range nums {
		pre += num
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre] += 1
	}
	return count
}
