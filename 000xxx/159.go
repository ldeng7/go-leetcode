func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	l, r, n := 0, -1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		if r >= 0 && s[i] != s[r] {
			n, l = max(n, i-l), r+1
		}
		r = i - 1
	}
	return max(n, len(s)-l)
}
