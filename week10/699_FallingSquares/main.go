package main

// O(N^2) time
// O(N) space
func fallingSquares(positions [][]int) []int {
	n := len(positions)
	qans := make([]int, n)
	for i, position := range positions {
		left := position[0]
		size := position[1]
		right := left + size
		qans[i] += size

		for j := i + 1; j < n; j++ {
			left2 := positions[j][0]
			size2 := positions[j][1]
			right2 := left2 + size2
			if left2 < right && left < right2 {
				qans[j] = max(qans[j], qans[i])
			}
		}
	}
	ans := []int{}
	cur := -1
	for _, x := range qans {
		cur = max(cur, x)
		ans = append(ans, cur)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
