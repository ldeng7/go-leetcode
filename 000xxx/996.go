import (
	"math"
	"sort"
)

func numSquarefulPerms(A []int) int {
	o, l := 0, len(A)
	t := make([]bool, l)
	sort.Ints(A)

	var cal func(int, int)
	cal = func(j int, c int) {
		if c == l {
			o += 1
			return
		}
		for i, a := range A {
			if t[i] {
				continue
			}
			b := false
			for k := i - 1; k >= 0; k-- {
				if !t[k] {
					if A[k] == a {
						b = true
					}
					break
				}
			}
			if b {
				continue
			}
			j += a
			fj := math.Sqrt(float64(j))
			if fj == float64(int(fj)) {
				t[i] = true
				cal(a, c+1)
				t[i] = false
			}
			j -= a
		}
	}

	for i, a := range A {
		if i > 0 && a == A[i-1] {
			continue
		}
		t[i] = true
		cal(a, 1)
		t[i] = false
	}
	return o
}
