package main

// O(m^2*n) time
// O(n) space
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ans := 0
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for c, v := range row {
				sum[c] += v
			}
			ans += subarraySum(sum, target)
		}
	}
	return ans
}

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for _, num := range nums {
		pre += num
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre]++
	}
	return count
}
