import "sort"

func cal(ar []int) []int {
	l := len(ar)
	o := make([]int, l)
	for i := 0; i < l; i++ {
		o[i] = -1
	}
	st := []int{}
	for _, a := range ar {
		for 0 != len(st) && a > st[len(st)-1] {
			t := st[len(st)-1]
			o[t] = a
			st = st[:len(st)-1]
		}
		st = append(st, a)
	}
	return o
}

func oddEvenJumps(A []int) int {
	l := len(A)
	indices := make([]int, l)
	for i := 0; i < l; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		a, b := indices[i], indices[j]
		return A[a] < A[b] || (A[a] == A[b] && a < b)
	})
	on := cal(indices)
	sort.Slice(indices, func(i, j int) bool {
		a, b := indices[i], indices[j]
		return A[a] > A[b] || (A[a] == A[b] && a < b)
	})
	en := cal(indices)

	o, e := make([]int, l), make([]int, l)
	o[l-1], e[l-1] = 1, 1
	for i := l - 2; i >= 0; i-- {
		if on[i] != -1 {
			o[i] = e[on[i]]
		}
		if en[i] != -1 {
			e[i] = o[en[i]]
		}
	}
	s := 0
	for _, i := range o {
		s += i
	}
	return s
}
