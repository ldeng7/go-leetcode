import (
	"math"
	"sort"
)

func smallestRange(nums [][]int) []int {
	var out []int
	ps, m := [][2]int{}, map[int]int{}
	for i, ns := range nums {
		for _, n := range ns {
			ps = append(ps, [2]int{n, i})
		}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i][0] < ps[j][0] || (ps[i][0] == ps[j][0] && ps[i][1] < ps[j][1])
	})
	l, c, d, n := 0, 0, math.MaxInt64, len(nums)
	for r, pr := range ps {
		if m[pr[1]] == 0 {
			c++
		}
		m[pr[1]]++
		for c == n && l <= r {
			pl := ps[l]
			nl, nr := pl[0], pr[0]
			if d > nr-nl {
				d = nr - nl
				out = []int{nl, nr}
			}
			m[pl[1]]--
			if m[pl[1]] == 0 {
				c--
			}
			l++
		}
	}
	return out
}
