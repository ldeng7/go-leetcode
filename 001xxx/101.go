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

func earliestAcq(logs [][]int, N int) int {
	sort.Slice(logs, func(i, j int) bool { return logs[i][0] < logs[j][0] })
	au := (&ArrayUnion{}).Init(N)
	for _, log := range logs {
		if au.Merge(log[2], log[1]) {
			N--
		}
		if 1 == N {
			return log[0]
		}
	}
	return -1
}
