import "sort"

type ArrayUnion struct {
	arr []int
}

func (au *ArrayUnion) Init(l int) *ArrayUnion {
	au.arr = make([]int, l)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
	return au
}

func (au *ArrayUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *ArrayUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == -1 || r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	au := (&ArrayUnion{}).Init(n + 1)
	o := 0
	for _, w := range wells {
		o += w
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})
	for _, p := range pipes {
		r1, r2 := au.GetRoot(p[0]-1), au.GetRoot(p[1]-1)
		if r1 == r2 {
			continue
		}
		w1, w2 := wells[r1], wells[r2]
		s1, s2 := w1+w2, min(w1, w2)+p[2]
		if s1 > s2 {
			o += s2 - s1
			if w1 > w2 {
				au.Set(r1, r2)
			} else {
				au.Set(r2, r1)
			}

		}
	}
	return o
}
