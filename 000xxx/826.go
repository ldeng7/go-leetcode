func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	t := make([]int, 100001)
	for i, d := range difficulty {
		if profit[i] > t[d] {
			t[d] = profit[i]
		}
	}
	for i := 1; i < 100001; i++ {
		if t[i-1] > t[i] {
			t[i] = t[i-1]
		}
	}
	o := 0
	for _, w := range worker {
		o += t[w]
	}
	return o
}
