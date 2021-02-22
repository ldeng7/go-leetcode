func frequencySort(nums []int) []int {
	l := len(nums)
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	ar := make([][2]int, 0, len(m))
	for n, c := range m {
		ar = append(ar, [2]int{c, n})
	}
	sort.Slice(ar, func(i, j int) bool {
		a, b := ar[i], ar[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	o := make([]int, 0, l)
	for _, p := range ar {
		for c := p[0]; c > 0; c-- {
			o = append(o, p[1])
		}
	}
	return o
}
