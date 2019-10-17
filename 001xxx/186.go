import "math"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maximumSum(arr []int) int {
	l := len(arr)
	t0, t1 := make([]int, l), make([]int, l)
	t0[0], t1[0] = arr[0], math.MinInt16
	for i := 1; i < l; i++ {
		t0[i] = max(t0[i-1]+arr[i], arr[i])
		t1[i] = max(t1[i-1]+arr[i], t0[i-1])
	}
	o := math.MinInt16
	for i := 0; i < l; i++ {
		o = max(o, max(t0[i], t1[i]))
	}
	return o
}
