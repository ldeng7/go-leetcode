import (
	"sort"
	"strconv"
)

func largestNumber(cost []int, target int) string {
	t := make([]string, target+1)
	for i := 1; i <= target; i++ {
		var s string
		for j, c := range cost {
			if p := i - c; p == 0 || (p > 0 && len(t[p]) > 0) {
				if s1 := t[p] + strconv.Itoa(j+1); len(s1) >= len(s) {
					s = s1
				}
			}
		}
		t[i] = s
	}

	s := t[target]
	if 0 == len(s) {
		return "0"
	}
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] > bs[j] })
	return string(bs)
}
