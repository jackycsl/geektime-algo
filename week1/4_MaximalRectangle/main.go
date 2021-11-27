package main

// Build histogram start from first row, increase heights if 1
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	maxArea := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

type Rect struct {
	witdh  int
	height int
}

func largestRectangleArea(heights []int) int {
	s := []Rect{}
	ans := 0
	heights = append(heights, 0)
	for _, height := range heights {
		accumulatedWidth := 0
		// previous height >= current height
		for len(s) != 0 && s[len(s)-1].height >= height {
			accumulatedWidth += s[len(s)-1].witdh
			ans = max(ans, s[len(s)-1].height*accumulatedWidth)
			s = s[:len(s)-1]
		}
		s = append(s, Rect{accumulatedWidth + 1, height})
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
