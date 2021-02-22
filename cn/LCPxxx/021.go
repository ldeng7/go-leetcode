func chaseGame(edges [][]int, startA int, startB int) int {
	l := len(edges)
	tbl := make([][]int, l+1)
	for _, e := range edges {
		i, j := e[0], e[1]
		if (i == startA && j == startB) || (i == startB && j == startA) {
			return 1
		}
		tbl[i] = append(tbl[i], j)
		tbl[j] = append(tbl[j], i)
	}

	ds := make([]int, l+1)
	ps := make([]int, l+1)
	bs := make([]bool, l+1)
	c := 0
	var dfs func(int, int)
	dfs = func(i, p int) {
		ds[i], ps[i] = ds[p]+1, p
		for _, j := range tbl[i] {
			if j == p {
				continue
			}
			if 0 == ds[j] {
				dfs(j, i)
			} else if ds[j] < ds[i] {
				for k := i; k != j; c++ {
					bs[k] = true
					k = ps[k]
				}
				bs[j] = true
				c++
			}
		}
	}
	dfs(1, 0)

	bfs := func(i int, b bool) []int {
		o := make([]int, l+1)
		for i := 0; i <= l; i++ {
			o[i] = 100001
		}
		q := make([]int, 1, 32)
		o[i], q[0] = 0, i
		for len(q) > 0 {
			t := q[0]
			q = q[1:]
			if b && bs[t] {
				return []int{t, o[t]}
			}
			for _, i := range tbl[t] {
				if o[i] <= o[t]+1 {
					continue
				}
				o[i] = o[t] + 1
				q = append(q, i)
			}
		}
		return o
	}

	ar1, ar2 := bfs(startA, false), bfs(startB, false)
	if c >= 4 {
		ar := bfs(startB, true)
		if ar[1]+1 < ar1[ar[0]] {
			return -1
		}
	}
	o := 0
	for i := 1; i <= l; i++ {
		if (ar1[i] > ar2[i]+1) && ar1[i] > o {
			o = ar1[i]
		}
	}
	return o
}
