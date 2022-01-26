package main

import (
	"fmt"
)

const max int = 5005

func main() {

	var n int
	fmt.Scan(&n)

	// a := make([]int, max)
	// sum := make([]int, max)
	// f := make([]int, max)
	// last := make([]int, max)

	a := make([]int, n+1)
	sum := make([]int, n+1)
	f := make([]int, n+1)
	last := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		sum[i] = sum[i-1] + a[i]
	}

	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if sum[i]-sum[j] >= last[j] {
				f[i] = f[j] + 1
				last[i] = sum[i] - sum[j]
			}
		}
	}
	fmt.Println(n - f[n])
}
