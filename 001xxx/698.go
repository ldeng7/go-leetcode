import "strings"

func countDistinct(s string) int {
	l := len(s)
	t := make([]int, l)
	t[0] = 1
	bs := [26]bool{}
	bs[s[0]-'a'] = true

	for i := 1; i < l; i++ {
		if j := s[i] - 'a'; !bs[j] {
			t[i] = t[i-1] + i + 1
			bs[j] = true
		} else {
			k := -1
			for j := i; j > -1; j-- {
				if !strings.Contains(s[:i], s[j:i+1]) {
					k = j + 1
					break
				}
			}
			t[i] = t[i-1] + k
		}
	}
	return t[l-1]
}
