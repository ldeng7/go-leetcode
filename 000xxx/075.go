func sortColors(nums []int) {
	c0, c1, c2 := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			c0++
		} else if n == 1 {
			c1++
		} else {
			c2++
		}
	}
	i := 0
	for ; i < c0; i++ {
		nums[i] = 0
	}
	e := i + c1
	for ; i < e; i++ {
		nums[i] = 1
	}
	e = i + c2
	for ; i < e; i++ {
		nums[i] = 2
	}
}
