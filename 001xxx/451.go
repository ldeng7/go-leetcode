import (
	"fmt"
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	ss := strings.Split(text, " ")
	s0 := []byte(ss[0])
	s0[0] = 'a' + s0[0] - 'A'
	ss[0] = string(s0)

	l := len(ss)
	ar := make([]int, l)
	for i := 0; i < l; i++ {
		ar[i] = i
	}
	sort.Slice(ar, func(i, j int) bool {
		a, b := ar[i], ar[j]
		la, lb := len(ss[a]), len(ss[b])
		return la < lb || (la == lb && a < b)
	})

	ss1 := make([]string, l)
	for i, p := range ar {
		ss1[i] = ss[p]
	}
	s0 = []byte(ss1[0])
	s0[0] = 'A' + s0[0] - 'a'
	ss1[0] = string(s0)
	return strings.Join(ss1, " ")
}
