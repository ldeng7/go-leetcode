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

func (au *ArrayUnion) Merge(i, j int) bool {
	r1, r2 := au.GetRoot(i), au.GetRoot(j)
	if r1 != r2 {
		au.arr[r1] = r2
		return true
	}
	return false
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	for i, q := range queries {
		queries[i] = append(q, i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	au := (&ArrayUnion{}).Init(n)
	o := make([]bool, len(queries))
	j := 0
	for _, q := range queries {
		for ; j < len(edgeList) && edgeList[j][2] < q[2]; j++ {
			au.Merge(edgeList[j][0], edgeList[j][1])
		}
		o[q[3]] = au.GetRoot(q[0]) == au.GetRoot(q[1])
	}
	return o
}
