import "strconv"

func generateAbbreviations(word string) []string {
	out := []string{}
	var cal func(string, int, int)
	cal = func(s string, i, c int) {
		if i != len(word) {
			cal(s, i+1, c+1)
			if c > 0 {
				s += strconv.Itoa(c)
			}
			s += word[i : i+1]
			cal(s, i+1, 0)
		} else {
			if c > 0 {
				s += strconv.Itoa(c)
			}
			out = append(out, s)
		}
	}
	cal("", 0, 0)
	return out
}
