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

func smallestEquivalentString(A string, B string, S string) string {
	au := (&ArrayUnion{}).Init(26)

	for i := 0; i < len(A); i++ {
		r1, r2 := au.GetRoot(int(A[i]-'a')), au.GetRoot(int(B[i]-'a'))
		if r1 > r2 {
			au.Set(r1, r2)
		} else {
			au.Set(r2, r1)
		}
	}
	bs := make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		bs[i] = byte(au.GetRoot(int(S[i]-'a'))) + 'a'
	}
	return string(bs)
}
