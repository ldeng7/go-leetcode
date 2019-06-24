import "sort"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestStrChain(words []string) int {
	o, l := 0, len(words)
	t := make([]int, l)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			a, b := words[i], words[j]
			if len(a) != len(b)+1 {
				continue
			}
			for k := 0; k < len(a); k++ {
				if a[:k]+a[k+1:] == b {
					t[i] = t[j] + 1
					o = max(o, t[i])
					break
				}
			}
		}
	}
	return o + 1
}
