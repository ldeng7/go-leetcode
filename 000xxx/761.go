import (
	"sort"
	"strings"
)

func makeLargestSpecial(S string) string {
	c, ar := 0, []string{}
	for i, j := 0, 0; i < len(S); i++ {
		if S[i] == '1' {
			c++
		} else {
			c--
		}
		if c == 0 {
			ar = append(ar, "1"+makeLargestSpecial(S[j+1:i])+"0")
			j = i + 1
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(ar)))
	return strings.Join(ar, "")
}
