import (
	"strconv"
)

func isAbbr(word, abbr string) bool {
	lw := len(word)
	j, c := 0, 0
	for _, b := range []byte(abbr) {
		if b >= '0' && b <= '9' {
			if b == '0' && c == 0 {
				return false
			}
			c = c*10 + int(b-'0')
		} else {
			j += c
			if j >= lw || b != word[j] {
				return false
			}
			j, c = j+1, 0
		}
	}
	return j+c == lw
}

func minAbbreviation(target string, dictionary []string) string {
	if 0 == len(dictionary) {
		return strconv.Itoa(len(target))
	}
	m := make([][]string, len(target))
	for i := 0; i < len(target)-1; i++ {
		m[i] = []string{}
	}

	var cal func(string, int, int)
	cal = func(s string, i, c int) {
		if i != len(target) {
			cal(s, i+1, c+1)
			if c > 0 {
				s += strconv.Itoa(c)
			}
			s += target[i : i+1]
			cal(s, i+1, 0)
		} else {
			if c > 0 {
				s += strconv.Itoa(c)
			}
			m[len(s)-1] = append(m[len(s)-1], s)
		}
	}
	cal("", 0, 0)

	for _, abbrs := range m {
		for _, abbr := range abbrs {
			i := 0
			for ; i < len(dictionary); i++ {
				if isAbbr(dictionary[i], abbr) {
					break
				}
			}
			if i == len(dictionary) {
				return abbr
			}
		}
	}
	return ""
}
