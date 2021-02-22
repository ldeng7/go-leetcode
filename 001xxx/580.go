import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	l := len(warehouse)
	arl, arr := make([]int, l), make([]int, l)
	arl[0] = warehouse[0]
	for i := 1; i < l; i++ {
		arl[i] = min(arl[i-1], warehouse[i])
	}
	arr[l-1] = warehouse[l-1]
	for i := l - 2; i >= 0; i-- {
		arr[i] = min(arr[i+1], warehouse[i])
	}
	for i := 0; i < l; i++ {
		warehouse[i] = max(arl[i], arr[i])
	}
	sort.Ints(warehouse)
	sort.Ints(boxes)
	o := 0
	for i, j := 0, 0; i < l && j < len(boxes); i++ {
		if boxes[j] <= warehouse[i] {
			o, j = o+1, j+1
		}
	}
	return o
}
