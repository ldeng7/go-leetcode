func deleteAndEarn(nums []int) int {
	ss := make([]int, 10001)
	for _, n := range nums {
		ss[n] += n
	}
	for i := 2; i < 10001; i++ {
		if ss[i-1] > ss[i]+ss[i-2] {
			ss[i] = ss[i-1]
		} else {
			ss[i] = ss[i] + ss[i-2]
		}
	}
	return ss[10000]
}
