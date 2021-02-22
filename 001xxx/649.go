const N = 100001

var bit [N]int
var cnt [N]int

func add(i int) {
	for ; i < N; i += (i & -i) {
		bit[i]++
	}
}

func sum(i int) int {
	o := 0
	for ; i > 0; i -= (i & -i) {
		o += bit[i]
	}
	return o
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func createSortedArray(instructions []int) int {
	o := 0
	for i := 0; i < N; i++ {
		bit[i], cnt[i] = 0, 0
	}
	for i, v := range instructions {
		l := sum(v - 1)
		r := i - l - cnt[v]
		o = (o + min(l, r)) % 1000000007
		cnt[v]++
		add(v)
	}
	return o
}
