func numRollsToTarget(d int, f int, target int) int {
	if target < 0 || target > d*f {
		return 0
	}
	t := []int{1}
	for i := 1; i <= d; i++ {
		t1 := make([]int, i*f+1)
		for j := 1; j <= f; j++ {
			for k, n := range t {
				t1[k+j] = (t1[k+j] + n) % 1000000007
			}
		}
		t = t1
	}
	return t[target]
}
