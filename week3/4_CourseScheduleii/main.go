package main

// O(n + m) time n is number of courses, m is prereq
// O(n + m) space
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	result := []int{}
	// course  0    1   2   3
	// edges [[1 2],[3],[3],[]]
	// indeg  [0 ,  1,  1,  2]
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		indeg[pre[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		// prereq
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(result) < numCourses {
		return nil
	}

	return result
}
