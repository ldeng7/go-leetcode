func findMaxConsecutiveOnes(nums []int) int {
	o, c, l, k := 0, 0, 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			c++
		}
		for c > k {
			if nums[l] == 0 {
				c--
			}
			l++
		}
		if i-l+1 > o {
			o = i - l + 1
		}
	}
	return o
}
