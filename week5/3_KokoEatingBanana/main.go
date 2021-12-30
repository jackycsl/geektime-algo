package main

import "math"

// O(N log W) time, n number of piles, W max of piles
// O(1) space
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	// [1, max(piles)]
	for _, w := range piles {
		if w > right {
			right = w
		}
	}

	for left < right {
		mid := (left + right) / 2

		if totalTime(mid, piles) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func totalTime(k int, piles []int) int {
	time := 0
	for _, p := range piles {
		time += int(math.Ceil(float64(p) / float64(k)))
		// time += (p - 1) / k + 1
	}
	return time
}
