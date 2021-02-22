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

func (au *ArrayUnion) Set(i, v int) {
	au.arr[i] = v
}

func (au *ArrayUnion) Reset() {
	l := len(au.arr)
	for i := 0; i < l; i++ {
		au.arr[i] = -1
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	r := make([]int, m+n)
	t := map[int][]uint32{}
	ar := make([]int, 0, 16)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v, k := matrix[i][j], uint32(i<<16)|uint32(j)
			if ks, ok := t[v]; ok {
				t[v] = append(ks, k)
			} else {
				ks := make([]uint32, 1, 16)
				ks[0], t[v] = k, ks
				ar = append(ar, v)
			}
		}
	}
	sort.Ints(ar)

	r1 := make([]int, m+n)
	au := (&ArrayUnion{}).Init(m + n)
	for _, v := range ar {
		copy(r1, r)
		au.Reset()
		for _, k := range t[v] {
			i, j := int(k>>16), int(k&0xffff)
			ri, rj := au.GetRoot(i), au.GetRoot(m+j)
			if ri != rj {
				rm := max(r1[ri], r1[rj])
				r1[ri], r1[rj] = rm, rm
				au.Set(ri, rj)
			}
		}

		for _, k := range t[v] {
			i, j := int(k>>16), int(k&0xffff)
			t := r1[au.GetRoot(i)] + 1
			r[i], r[m+j], matrix[i][j] = t, t, t
		}
	}
	return matrix
}
