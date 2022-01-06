package main

// O(mn * α mn) time
// O(mn) space
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	var dx = []int{-1, 0, 0, 1}
	var dy = []int{0, -1, 1, 0}

	fa := make([]int, m*n+1)
	for i := range fa {
		fa[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if x == fa[x] {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}

	var unionSet func(x, y int)
	unionSet = func(x, y int) {
		x, y = find(x), find(y)
		if x != y {
			fa[x] = y
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < 4; k++ {
				ni := i + dx[k]
				nj := j + dy[k]

				if ni < 0 || nj < 0 || ni >= m || nj >= n {
					continue
				} else if grid[ni][nj] == '1' {
					unionSet(nums(i, j, n), nums(ni, nj, n))
				}
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && find(nums(i, j, n)) == nums(i, j, n) {
				ans++
			}
		}
	}
	return ans
}

func nums(i, j, n int) int {
	return i*n + j
}
