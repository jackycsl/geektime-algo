package main

// Time O(n log n) worst, average O(n αn)
// Space O(n)
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
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

	for _, e := range edges {
		if find(e[0]) != find(e[1]) {
			unionSet(e[0], e[1])
		} else {
			return e
		}
	}
	return nil
}
