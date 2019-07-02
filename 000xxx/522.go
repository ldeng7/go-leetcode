import "sort"

func findLUSlength(strs []string) int {
	l := len(strs)
	s := map[string]bool{}
	sort.Slice(strs, func(i, j int) bool {
		a, b := strs[i], strs[j]
		if len(a) == len(b) {
			return a > b
		}
		return len(a) > len(b)
	})
	for i, str := range strs {
		if i == l-1 || str != strs[i+1] {
			b := true
		outer:
			for str1, _ := range s {
				j := 0
				for k := 0; k < len(str1); k++ {
					if str1[k] == str[j] {
						j++
					}
					if j == len(str) {
						b = false
						break outer
					}
				}
			}
			if b {
				return len(str)
			}
		}
		s[str] = true
	}
	return -1
}
