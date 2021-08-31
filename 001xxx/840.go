import "sort"

func min(ar, b int) int {
	if ar <= b {
		return ar
	}
	return b
}

func max(ar, b int) int {
	if ar >= b {
		return ar
	}
	return b
}

func abs(ar int) int {
	if ar >= 0 {
		return ar
	}
	return -ar
}

func maxBuilding(n int, restrictions [][]int) int {
	ar := append(restrictions, []int{1, 0})
	l := len(ar)
	sort.Slice(ar, func(i, j int) bool {
		return ar[i][0] < ar[j][0]
	})
	for i := 1; i < l; i++ {
		ar[i][1] = min(ar[i][1], ar[i-1][1]+ar[i][0]-ar[i-1][0])
	}
	for i := l - 2; i >= 0; i-- {
		ar[i][1] = min(ar[i][1], ar[i+1][1]+ar[i+1][0]-ar[i][0])
	}

	o := 0
	for i := 1; i < l; i++ {
		o = max(o, max(ar[i-1][1], ar[i][1])+(ar[i][0]-ar[i-1][0]-abs(ar[i-1][1]-ar[i][1]))/2)
	}
	return max(o, ar[l-1][1]+n-ar[l-1][0])
}
