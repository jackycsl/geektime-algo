package main

// O(N) time: go through the length of all num
// O(N) space: Hashmap length of all num

type entry struct{ count, left, right int }

func findShortestSubArray(nums []int) int {
	mp := make(map[int]entry)
	for i, num := range nums {
		if e, ok := mp[num]; ok {
			e.count++
			e.right = i
			mp[num] = e
		} else {
			mp[num] = entry{1, i, i}
		}
	}
	maxNum := 0
	minLen := 0
	for _, e := range mp {
		if e.count > maxNum {
			maxNum = e.count
			minLen = e.right - e.left + 1
		} else if e.count == maxNum {
			minLen = min(minLen, e.right-e.left+1)
		}
	}
	return minLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
