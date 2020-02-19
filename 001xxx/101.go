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
	for {
		r := au.arr[i]
		if r == i || r == -1 {
			return i
		}
		i = r
	}
}

func (au *ArrayUnion) Union(a, b int) bool {
	ra, rb := au.GetRoot(a), au.GetRoot(b)
	if ra == rb {
		return false
	}
	au.arr[ra] = rb
	return true
}

func earliestAcq(logs [][]int, N int) int {
	sort.Slice(logs, func(i, j int) bool { return logs[i][0] < logs[j][0] })
	au := (&ArrayUnion{}).Init(N)
	for _, log := range logs {
		if au.Union(log[2], log[1]) {
			N--
		}
		if 1 == N {
			return log[0]
		}
	}
	return -1
}
