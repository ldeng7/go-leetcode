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
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func (au *ArrayUnion) Reset() {
	l := len(au.arr)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	au := (&ArrayUnion{}).Init(n)
	cal := func(b, m int) int {
		au.Reset()
		o := 0
		if m != -1 {
			e := edges[m]
			au.Set(au.GetRoot(e[1]), au.GetRoot(e[0]))
			o += e[2]
		}
		for i, e := range edges {
			if i == b {
				continue
			}
			if r1, r2 := au.GetRoot(e[0]), au.GetRoot(e[1]); r1 != r2 {
				au.Set(r2, r1)
				o += e[2]
			}
		}
		return o
	}

	for i, e := range edges {
		edges[i] = append(e, i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	b := cal(-1, -1)
	o1, o2 := make([]int, 0, len(edges)), make([]int, 0, len(edges))
	for i, e := range edges {
		if b != cal(i, -1) {
			o1 = append(o1, e[3])
		} else if b == cal(-1, i) {
			o2 = append(o2, e[3])
		}
	}
	return [][]int{o1, o2}
}
