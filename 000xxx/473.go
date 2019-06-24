func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	s := 0
	for _, n := range nums {
		s += n
	}
	if s&3 != 0 {
		return false
	}

	l := uint(len(nums))
	all, t := (1<<l)-1, s>>2
	masks, bs := []int{}, make([]bool, 1<<l)
	for i := 0; i <= all; i++ {
		s := 0
		for j := uint(0); j <= 15; j++ {
			if (i>>j)&1 == 1 {
				s += nums[j]
			}
		}
		if s == t {
			for _, m := range masks {
				if (m & i) != 0 {
					continue
				}
				h := m | i
				bs[h] = true
				if bs[all^h] {
					return true
				}
			}
			masks = append(masks, i)
		}
	}
	return false
}
