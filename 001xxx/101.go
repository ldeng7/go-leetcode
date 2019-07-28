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

func earliestAcq(logs [][]int, N int) int {
	sort.Slice(logs, func(i, j int) bool { return logs[i][0] < logs[j][0] })
	au := (&ArrayUnion{}).Init(N)
	for _, log := range logs {
		r1, r2 := au.GetRoot(log[1]), au.GetRoot(log[2])
		if r1 != r2 {
			au.Set(r2, r1)
			N--
		}
		if 1 == N {
			return log[0]
		}
	}
	return -1
}
