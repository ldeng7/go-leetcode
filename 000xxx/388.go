import "strings"

func lengthLongestPath(input string) int {
	o, m := 0, map[int]int{}
	m[0] = 0
	ss := strings.Split(input, "\n")
	for _, s := range ss {
		l := strings.LastIndexByte(s, '\t') + 1
		n := len(s) - l
		if -1 != strings.IndexByte(s, '.') {
			if o1 := m[l] + n; o1 > o {
				o = o1
			}
		} else {
			m[l+1] = m[l] + n + 1
		}
	}
	return o
}
