import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDeletions(s string) int {
	l := len(s)
	cnts := [26]int{}
	for i := 0; i < l; i++ {
		cnts[s[i]-'a']++
	}
	sort.Slice(cnts[:], func(i, j int) bool {
		return cnts[i] > cnts[j]
	})
	n, p := 0, 100001
	for _, c := range cnts {
		if c == 0 || p == 0 {
			break
		}
		p = min(c, p-1)
		n += p
	}
	return l - n
}
