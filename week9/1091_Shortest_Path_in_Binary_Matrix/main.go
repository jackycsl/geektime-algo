package main

// O(m + n) time
// O(m + n) space
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 1
	}

	dx := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}

	q := make([]int, 0)
	q = append(q, 0)
	grid[0][0] = 1

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		x, y := cur/n, cur%n

		for i := 0; i < 8; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				q = append(q, nx*n+ny)
				grid[nx][ny] = grid[x][y] + 1
				if nx == m-1 && ny == n-1 {
					return grid[nx][ny]
				}
			}
		}
	}
	return -1
}
