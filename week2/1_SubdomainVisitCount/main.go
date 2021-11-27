package main

import (
	"strconv"
	"strings"
)

// O(N) time: go through the length of cpdomains
// O(N) space: HashMap
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	ans := make([]string, 0, len(cpdomains))

	for _, cpdomain := range cpdomains {
		d := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(d[0])
		domain := d[1]
		frags := strings.Split(domain, ".")
		for i := range frags {
			m[strings.Join(frags[i:], ".")] += count
		}
	}

	for k, v := range m {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}

	return ans
}
