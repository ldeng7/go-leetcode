import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	o, m, n := 0, len(boxes), len(warehouse)
	for i := 1; i < n; i++ {
		warehouse[i] = min(warehouse[i], warehouse[i-1])
	}
	sort.Ints(boxes)
	for i := 0; i < n && o < m; i++ {
		if warehouse[n-i-1] >= boxes[o] {
			o++
		}
	}
	return o
}
