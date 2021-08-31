import (
	"math"
	"sort"
)

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]int, n+1)
	bs := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		g[i] = make([]int, 0, 16)
		bs[i] = make([]bool, n+1)
	}
	for _, e := range edges {
		a, b := e[0], e[1]
		bs[a][b], bs[b][a] = true, true
		g[a], g[b] = append(g[a], b), append(g[b], a)
	}

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		ar[i] = i
	}
	sort.Slice(ar, func(i, j int) bool {
		return len(g[ar[i]]) < len(g[ar[j]])
	})
	o := math.MaxInt64

	for _, i := range ar {
		ar1 := g[i]
		l1 := len(ar1)
		if l1 < 2 {
			continue
		} else if l1 >= o-4 {
			break
		}
		sort.Slice(ar1, func(i, j int) bool {
			return len(g[ar1[i]]) < len(g[ar1[j]])
		})
		for j := 0; j < l1-1; j++ {
			a := ar1[j]
			l2 := len(g[a])
			if l1+l2+len(g[ar1[j+1]]) >= o {
				break
			}
			for k := j + 1; k < l1; k++ {
				if d := l1 + l2 + len(g[ar1[k]]); d >= o {
					break
				} else if bs[a][ar1[k]] {
					o = d
				}
			}
		}
	}
	if o == math.MaxInt64 {
		return -1
	}
	return o - 6
}
