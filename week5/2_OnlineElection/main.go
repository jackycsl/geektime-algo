package main

// O(n) time
// O(n) space
type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tops := []int{}
	top := -1
	count := make(map[int]int)
	for _, p := range persons {
		count[p]++
		if count[p] >= count[top] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{tops, times}
}

func (this *TopVotedCandidate) Q(t int) int {
	left, right := 0, len(this.times)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if this.times[mid] <= t {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return this.tops[left]
}

/**
* Your TopVotedCandidate object will be instantiated and called as such:
* obj := Constructor(persons, times);
* param_1 := obj.Q(t);
 */
