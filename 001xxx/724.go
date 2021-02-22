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

type DistanceLimitedPathsExist struct {
	au *ArrayUnion
	g  []map[int]int
	ds []int
	ps []int
	ls []int
}

func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
	this := &DistanceLimitedPathsExist{
		au: (&ArrayUnion{}).Init(n),
		g:  make([]map[int]int, n),
		ds: make([]int, n),
		ps: make([]int, n),
		ls: make([]int, n),
	}
	for i := 0; i < n; i++ {
		this.g[i] = map[int]int{}
	}
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	for _, e := range edgeList {
		a, b, w := e[0], e[1], e[2]
		if this.au.Merge(a, b) {
			this.g[a][b], this.g[b][a] = w, w
		}
	}

	v := make([]bool, n)
	var cal func(int)
	cal = func(i int) {
		v[i] = true
		for j, w := range this.g[i] {
			if v[j] {
				continue
			}
			this.ds[j], this.ps[j], this.ls[j] = this.ds[i]+1, i, w
			cal(j)
		}
	}
	for i := 0; i < n; i++ {
		if !v[i] {
			this.ps[i] = i
			cal(i)
		}
	}
	return *this
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
	if this.au.GetRoot(p) != this.au.GetRoot(q) {
		return false
	}
	l, dp, dq := 0, this.ds[p], this.ds[q]
	if dp < dq {
		p, q = q, p
		dp, dq = dq, dp
	}
	for i := 0; i < dp-dq; i++ {
		l = max(l, this.ls[p])
		p = this.ps[p]
	}
	for p != q {
		l = max(l, max(this.ls[p], this.ls[q]))
		p, q = this.ps[p], this.ps[q]
	}
	return l < limit
}
