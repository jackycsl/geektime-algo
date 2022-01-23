package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *hp) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}

// O(n log n) time
// O(n) space
func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}

	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]

	for i := k; i < n; i++ {
		heap.Push(q, i)
		fmt.Println(q, i, k)
		// [i-k+1, i] max value
		// 2 3 [-1 -3 1]
		// pop 2 and 3
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
