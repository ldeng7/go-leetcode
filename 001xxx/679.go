func maxOperations(nums []int, k int) int {
	m := map[int]int{}
	o := 0
	for _, v := range nums {
		if v > k {
			continue
		}
		if m[k-v] != 0 {
			m[k-v]--
			o++
		} else {
			m[v]++
		}
	}
	return o
}
