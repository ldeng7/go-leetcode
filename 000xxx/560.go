func subarraySum(nums []int, k int) int {
	out, sum := 0, 0
	m := map[int]int{0: 1}
	for _, n := range nums {
		sum += n
		out += m[sum-k]
		m[sum] = m[sum] + 1
	}
	return out
}
