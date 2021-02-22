import "sort"

func canDistribute(nums []int, quantity []int) bool {
	cnts := map[int]int{}
	for _, n := range nums {
		cnts[n]++
	}
	arc := make([]int, 0, len(cnts))
	for _, c := range cnts {
		arc = append(arc, c)
	}
	sort.Slice(arc, func(i, j int) bool { return arc[i] > arc[j] })
	sort.Slice(quantity, func(i, j int) bool { return quantity[i] > quantity[j] })

	m := (1 << len(quantity)) - 1
	mqs := make([]int, m+1)
	for i := 1; i <= m; i++ {
		for j, q := range quantity {
			if (i>>j)&1 == 1 {
				mqs[i] += q
			}
		}
	}

	m1 := 0
	for i := 0; i < len(arc) && m1 != m; i++ {
		mmq, mm := 0, 0
		for j := 1; j <= m; j++ {
			if j&m1 == 0 && mqs[j] <= arc[i] && mqs[j] > mmq {
				mmq, mm = mqs[j], j
			}
		}
		m1 |= mm
		if mm == 0 {
			break
		}
	}
	return (m1 == m)
}
