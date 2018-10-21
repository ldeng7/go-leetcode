func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] = m[num] + 1
	}
	h := len(nums) >> 1
	if len(nums)&1 == 1 {
		h++
	}
	for num, cnt := range m {
		if cnt >= h {
			return num
		}
	}
	return 0
}
