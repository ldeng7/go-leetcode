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

func (au *ArrayUnion) GetRoot(i int) int {
	r := au.arr[i]
	if r == -1 || r == i {
		return i
	}
	r = au.GetRoot(r)
	au.arr[i] = r
	return r
}

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func minimumCost(N int, connections [][]int) int {
	au := (&ArrayUnion{}).Init(N + 1)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	o, c := 0, 0
	for _, conn := range connections {
		if au.Merge(conn[0], conn[1]) {
			o, c = o+conn[2], c+1
			if c == N-1 {
				return o
			}
		}
	}
	return -1
}
