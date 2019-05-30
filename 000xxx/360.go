func cal(x, a, b, c int) int {
	return a*x*x + b*x + c
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	n := len(nums)
	i, j, k := 0, n-1, 0
	if a >= 0 {
		k = n - 1
	}
	out := make([]int, n)
	for i <= j {
		oi, oj := cal(nums[i], a, b, c), cal(nums[j], a, b, c)
		if a >= 0 {
			if oi >= oj {
				out[k] = oi
				i++
			} else {
				out[k] = oj
				j--
			}
			k--
		} else {
			if oi >= oj {
				out[k] = oj
				j--
			} else {
				out[k] = oi
				i++
			}
			k++
		}
	}
	return out
}
