func abs(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -a
}

func judgePoint24(nums []int) bool {
	ar := make([]float64, len(nums))
	for i, n := range nums {
		ar[i] = float64(n)
	}

	var cal func([]float64) bool
	cal = func(nums []float64) bool {
		l := len(nums)
		if l == 1 {
			return abs(nums[0]-24) < 0.001
		}
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if i == j {
					continue
				}
				a, b := nums[i], nums[j]
				ar := []float64{}
				for k := 0; k < l; k++ {
					if k != i && k != j {
						ar = append(ar, nums[k])
					}
				}
				if i <= j {
					ar = append(ar, a+b)
					if cal(ar) {
						return true
					}
					ar = ar[:len(ar)-1]
				}
				ar = append(ar, a-b)
				if cal(ar) {
					return true
				}
				ar = ar[:len(ar)-1]
				if i <= j {
					ar = append(ar, a*b)
					if cal(ar) {
						return true
					}
					ar = ar[:len(ar)-1]
				}
				if b >= 0.001 {
					ar = append(ar, a/b)
					if cal(ar) {
						return true
					}
					ar = ar[:len(ar)-1]
				}
			}
		}
		return false
	}
	return cal(ar)
}
