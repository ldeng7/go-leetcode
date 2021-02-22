import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxValue(events [][]int, k int) int {
	l := len(events)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		ar := make([]int, k+1)
		for j := 0; j <= k; j++ {
			ar[j] = -1
		}
		t[i] = ar
	}
	sort.Slice(events, func(i, j int) bool {
		a, b := events[i], events[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		} else if a[1] != b[1] {
			return a[1] < b[1]
		}
		return a[2] < b[2]
	})

	var cal func(int, int) int
	cal = func(i, k int) int {
		if k == 0 || i >= l {
			return 0
		} else if v := t[i][k]; v != -1 {
			return v
		}
		v := events[i][1]
		j := i + sort.Search(l-i, func(k int) bool {
			return events[k+i][0] > v
		})
		v = max(events[i][2]+cal(j, k-1), cal(i+1, k))
		t[i][k] = v
		return v
	}
	return cal(0, k)
}
