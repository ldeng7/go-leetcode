import "sort"

func topKFrequent(words []string, k int) []string {
	l := len(words)
	m := map[string]int{}
	ss := make([]map[string]bool, l)
	for i := 0; i < l; i++ {
		ss[i] = map[string]bool{}
	}
	for _, w := range words {
		m[w]++
	}
	for w, c := range m {
		ss[c-1][w] = true
	}
	out := []string{}
	for i := l - 1; i >= 0; i-- {
		if k <= 0 {
			break
		}
		s, n := ss[i], len(ss[i])
		if n == 0 {
			continue
		}
		ar, j := make([]string, n), 0
		for w, _ := range s {
			ar[j], j = w, j+1
		}
		sort.Strings(ar)
		if k < len(ar) {
			ar = ar[:k]
		}
		out = append(out, ar...)
		k -= len(ar)
	}
	return out
}
