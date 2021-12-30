package main

// O(n log sum W) time
// O(1) space
func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	// [max(w), sum(w)]
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}

	for left < right {
		mid := (left + right) / 2
		d := countDays(mid, weights)

		if d <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func countDays(targetWeight int, weights []int) int {
	days := 1
	current := 0

	for _, weight := range weights {
		current += weight
		if current > targetWeight {
			days += 1
			current = weight
		}
	}
	return days
}
