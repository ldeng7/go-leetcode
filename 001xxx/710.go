import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	o, m := 0, 0
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		m = min(truckSize, boxTypes[i][0])
		truckSize -= m
		o += m * boxTypes[i][1]
	}
	return o
}
