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

func minimumCost(N int, connections [][]int) int {
	au := (&ArrayUnion{}).Init(N + 1)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	o, c := 0, 0
	for _, conn := range connections {
		r1, r2 := au.GetRoot(conn[0]), au.GetRoot(conn[1])
		if r1 != r2 {
			o, c = o+conn[2], c+1
			if c == N-1 {
				return o
			}
			au.Set(r1, r2)
		}
	}
	return -1
}
