import "sort"

var ar [1001][2]int

func init() {
	var cal func(int) int
	cal = func(i int) int {
		if i <= 1000 && ar[i][1] != 0 {
			return ar[i][0]
		} else if i == 1 {
			return 0
		} else if i&1 == 1 {
			return cal(i*3+1) + 1
		}
		return cal(i/2) + 1
	}
	for i := 1; i <= 1000; i++ {
		ar[i] = [2]int{cal(i), i}
	}
}

func getKth(lo int, hi int, k int) int {
	pairs := make([][2]int, hi-lo+1)
	copy(pairs, ar[lo:hi+1])
	sort.Slice(pairs, func(i, j int) bool {
		w1, w2 := pairs[i][0], pairs[j][0]
		return w1 < w2 || (w1 == w2 && pairs[i][1] < pairs[j][1])
	})
	return pairs[k-1][1]
}
