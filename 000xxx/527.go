import "fmt"

func abbr(s string, k int) string {
	if k >= len(s)-2 {
		return s
	}
	return fmt.Sprintf("%s%d%c", s[:k], len(s)-k-1, s[len(s)-1])
}

func wordsAbbreviation(dict []string) []string {
	n := len(dict)
	pre := make([]int, n)
	out := make([]string, n)
	for i := 0; i < n; i++ {
		pre[i], out[i] = 1, abbr(dict[i], 1)
	}
	for i := 0; i < n; i++ {
		for {
			s := map[int]bool{}
			for j := i + 1; j < n; j++ {
				if out[j] == out[i] {
					s[j] = true
				}
			}
			if 0 == len(s) {
				break
			}
			s[i] = true
			for k, _ := range s {
				pre[k]++
				out[k] = abbr(dict[k], pre[k])
			}
		}
	}
	return out
}
