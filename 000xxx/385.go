import (
	"strconv"
)

func deserialize(s string) *NestedInteger {
	l := len(s)
	if 0 == len(s) {
		return &NestedInteger{}
	} else if s[0] != '[' {
		ni := &NestedInteger{}
		i, _ := strconv.Atoi(s)
		ni.SetInteger(i)
		return ni
	} else if l <= 2 {
		return &NestedInteger{}
	}
	ni := &NestedInteger{}
	b, c := 1, 0
	for i := 1; i < l; i++ {
		if c == 0 && (s[i] == ',' || i == l-1) {
			ni.Add(*deserialize(s[b:i]))
			b = i + 1
		} else if s[i] == '[' {
			c++
		} else if s[i] == ']' {
			c--
		}
	}
	return ni
}
