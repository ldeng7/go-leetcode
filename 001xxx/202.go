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

func smallestStringWithSwaps(s string, pairs [][]int) string {
	o, l := []byte(s), len(s)
	au := (&ArrayUnion{}).Init(l)
	for _, p := range pairs {
		au.Merge(p[0], p[1])
	}
	g := map[int][]int{}
	for i := 0; i < l; i++ {
		r := au.GetRoot(i)
		g[r] = append(g[r], i)
	}
	for _, ar := range g {
		bs := make([]byte, 0, len(ar))
		for _, i := range ar {
			bs = append(bs, s[i])
		}
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		for i, c := range bs {
			o[ar[i]] = c
		}
	}
	return string(o)
}
