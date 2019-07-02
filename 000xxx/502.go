import "sort"

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	ar := [][2]int{}
	for i, c := range Capital {
		ar = append(ar, [2]int{c, Profits[i]})
	}
	sort.Slice(ar, func(i, j int) bool {
		a, b := ar[i], ar[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	for i := 0; i < k; i++ {
		l, r, mx, k := 0, len(ar), 0, 0
		for l < r {
			m := l + (r-l)>>1
			if ar[m][0] <= W {
				l = m + 1
			} else {
				r = m
			}
		}
		for j := r - 1; j >= 0; j-- {
			if t := ar[j][1]; t > mx {
				mx, k = t, j
			}
		}
		W += mx
		if k < len(ar) {
			copy(ar[k:], ar[k+1:])
			ar = ar[:len(ar)-1]
		}
	}
	return W
}
