package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var n int
	var m int
	var x int
	var y int

	fmt.Scan(&n)
	fmt.Scan(&m)

	mice := make([][]int, n)
	for i := range mice {
		mice[i] = make([]int, 2)
		fmt.Scan(&x)
		fmt.Scan(&y)
		mice[i][0] = x
		mice[i][1] = y
	}

	flower := make([][]int, m)
	for i := range flower {
		flower[i] = make([]int, 2)
		fmt.Scan(&x)
		fmt.Scan(&y)
		flower[i][0] = x
		flower[i][1] = y
	}

	// n := 11
	// m := 1
	// mice := make([][]int, n)
	// flower := make([][]int, m)
	// mice = [][]int{{10, 3}, {10, 3}, {2, 2}, {4, 1}, {4, 1}, {1, 4}, {1, 4}, {2, 3}, {2, 3}, {1, 1}, {1, 1}}
	// mice = [][]int{{1, 2}, {1, 2}, {2, 2}, {1, 2}}
	// flower = [][]int{{2, 2}}

	miceCounter := make(map[string]int)
	for i := 0; i < n; i++ {
		key := strconv.Itoa(mice[i][0]) + "_" + strconv.Itoa(mice[i][1])
		miceCounter[key] += 1
	}

	flowerSet := make(map[string]bool)
	for i := 0; i < m; i++ {
		key := strconv.Itoa(flower[i][0]) + "_" + strconv.Itoa(flower[i][1])
		if _, ok := flowerSet[key]; !ok {
			flowerSet[key] = true
		}
	}

	pq := PriorityQueue{}
	heap.Init(&pq)

	for key, count := range miceCounter {
		if _, ok := flowerSet[key]; ok {
			continue
		}
		item := &Item{
			key:   key,
			count: count,
		}
		heap.Push(&pq, item)
	}

	count := pq[0].count
	ansq := [][]int{}
	sum := math.MaxInt32

	for len(pq) != 0 {
		item := heap.Pop(&pq).(*Item)
		if item.count != count {
			break
		}
		s := strings.Split(item.key, "_")
		i, _ := strconv.Atoi(s[0])
		j, _ := strconv.Atoi(s[1])
		if i+j < sum {
			ansq = [][]int{}
			ansq = append(ansq, []int{i, j})
			sum = i + j
		} else if i+j == sum {
			ansq = append(ansq, []int{i, j})
			sum = i + j
		}
	}

	if len(ansq) < 2 {
		fmt.Println(ansq[0][0], ansq[0][1])
	} else {
		min := ansq[0][0]
		for i := 1; i < len(ansq); i++ {
			if ansq[i][0] < min {
				min = ansq[i][0]
			}
		}
		fmt.Println(min)
	}

}

type Item struct {
	key   string
	count int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
