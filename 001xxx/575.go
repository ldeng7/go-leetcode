import "sort"

const MOD = 1000000007

func countRoutes(locations []int, start int, finish int, fuel int) int {
	l := len(locations)
	start, finish = locations[start], locations[finish]
	sort.Ints(locations)
	s := sort.Search(l, func(j int) bool { return locations[j] >= start })
	f := sort.Search(l, func(j int) bool { return locations[j] >= finish })

	t, u := make([][]int, l), make([][]int, l)
	for j := 0; j < l; j++ {
		t[j], u[j] = make([]int, fuel+1), make([]int, fuel+1)
	}
	o, b := 0, true
	if s == f {
		o = 1
	}
	for i := fuel; i > 0 && b; i-- {
		b = false
		for j := 1; j < l; j++ {
			d := locations[j] - locations[j-1]
			if i < d {
				continue
			}
			b = true
			v := t[j-1][i]*2 + 1
			if j-1 != s || i != fuel {
				v += u[j-1][i] - 1
			}
			t[j][i-d] = v % MOD
			v = u[j][i]*2 + 1
			if j != s || i != fuel {
				v += t[j][i] - 1
			}
			u[j-1][i-d] = v % MOD
		}
		o = (o + t[f][i] + u[f][i]) % MOD
	}
	return (o + t[f][0] + u[f][0]) % MOD
}
