func pathSum(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	m := map[int]int{}
	out := 0
	for _, num := range nums {
		m[num/10] = num % 10
	}

	var cal func(num, i int)
	cal = func(num, i int) {
		l := (num/10+1)*10 + 2*(num%10) - 1
		r := l + 1
		i += m[num]
		_, le := m[l]
		_, re := m[r]
		if !le && !re {
			out += i
			return
		}
		if le {
			cal(l, i)
		}
		if re {
			cal(r, i)
		}
	}
	cal(nums[0]/10, 0)
	return out
}
