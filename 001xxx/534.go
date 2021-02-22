type FenwickTree struct {
	ar []int
}

func (t *FenwickTree) Init(n int) *FenwickTree {
	t.ar = make([]int, n+1)
	return t
}

func (t *FenwickTree) sum(i int) int {
	o := 0
	for j := i + 1; j > 0; j -= j & (-j) {
		o += t.ar[j]
	}
	return o
}

func (t *FenwickTree) Sum(i, j int) int {
	if j >= i {
		if i != 0 {
			return t.sum(j) - t.sum(i-1)
		}
		return t.sum(j)
	}
	return -1
}

func (t *FenwickTree) Add(i, a int) {
	for j := i + 1; j < len(t.ar); j += j & (-j) {
		t.ar[j] += a
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	l, m := len(arr), 0
	for _, a := range arr {
		m = max(m, a)
	}
	cs := make([]int, m+1)
	cs[arr[0]] = 1
	t := (&FenwickTree{}).Init(m + 1)
	for i := 2; i < l; i++ {
		t.Add(arr[i], 1)
	}
	o := 0
	for i := 1; i < l-1; i++ {
		n := arr[i]
		mia, maa, mib, mab := min(n+a, m), max(n-a, 0), min(n+b, m), max(n-b, 0)
		for j := maa; j <= mia; j++ {
			if cs[j] == 0 {
				continue
			}
			mi, ma := min(mib, min(j+c, m)), max(mab, max(j-c, 0))
			if mi >= ma {
				o += cs[j] * t.Sum(ma, mi)
			}
		}
		cs[n]++
		t.Add(arr[i+1], -1)
	}
	return o
}
