var facs = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101,
	103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197,
	199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313,
}

type Union struct {
	arr []int
}

func (au *Union) union(a, b int) {
	r1, r2 := au.find(a), au.find(b)
	if r1 != r2 {
		au.arr[r1] += au.arr[r2]
		au.arr[r2] = r1
	}
}
func (au *Union) find(a int) int {
	i := au.arr[a]
	if i <= 0 {
		return a
	}
	j := au.find(i)
	au.arr[a] = j
	return j
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func largestComponentSize(A []int) int {
	m := map[int]int{}
	u := &Union{make([]int, len(facs))}
	for _, a := range A {
		e := -1
		for i := 0; i < len(facs) && facs[i] <= a; i++ {
			f := facs[i]
			if a%f == 0 {
				for {
					a /= f
					if a%f != 0 {
						break
					}
				}
				if e >= 0 {
					u.union(i, e)
				}
				e = i
			}
		}

		if a > 1 {
			i, ok := m[a]
			if ok {
				if -1 != e && -1 != i {
					u.union(e, i)
				} else if -1 != e || -1 != i {
					u.arr[u.find(max(e, i))]--
				}
			}
			if !ok || e >= 0 {
				m[a] = e
			}
		}
		if e >= 0 {
			u.arr[u.find(e)]--
		}
	}

	o := 314
	for _, r := range u.arr {
		if r < o {
			o = r
		}
	}
	return -o
}
