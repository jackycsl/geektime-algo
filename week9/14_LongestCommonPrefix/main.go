package main

// O(mnlogm) time
// O(1) space
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var isCommonPrefix func(length int) bool

	isCommonPrefix = func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}

	left, right := 0, minLength
	for left < right {
		mid := (left + right + 1) >> 1
		if isCommonPrefix(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return strs[0][:left]
}
