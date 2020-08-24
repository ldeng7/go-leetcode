func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

var t [1 << 15]int

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	m := 1 << n
	t[0] = 0
	for i := 1; i < m; i++ {
		t[i] = -1
	}
	mask := make([]uint, n)
	for _, d := range dependencies {
		a, b := uint(d[0]-1), uint(d[1]-1)
		mask[b] |= (1 << a)
	}

	var ar []int
	var l int
	var cal func(int, int, int, int)
	cal = func(i, j, k, v int) {
		if l-i < k {
			return
		}
		if i >= l || k == 0 {
			if t[j] == -1 || t[j] > v {
				t[j] = v
			}
		} else {
			cal(i+1, j|(1<<ar[i]), k-1, v)
			cal(i+1, j, k, v)
		}
	}

	for i := 0; i < m; i++ {
		if t[i] == -1 {
			continue
		}
		ar = make([]int, 0, n)
		for j := 0; j < n; j++ {
			if (((i >> j) & 1) == 1) || ((uint(i) & mask[j]) != mask[j]) {
				continue
			}
			ar = append(ar, j)
		}
		l = len(ar)
		cal(0, i, min(l, k), t[i]+1)
	}
	return t[m-1]
}
