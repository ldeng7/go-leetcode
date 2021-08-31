func countNicePairs(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		r := 0
		for n := num; n > 0; n /= 10 {
			r = r*10 + n%10
		}
		m[num-r]++
	}
	o := 0
	for _, c := range m {
		o += (c * (c - 1)) / 2
	}
	return o % 1000000007
}
