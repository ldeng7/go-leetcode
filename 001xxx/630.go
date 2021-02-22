func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	o := make([]bool, n)
	for i := 0; i < n; i++ {
		m, mi, ma := r[i]-l[i]+1, 100000, -100000
		for j := l[i]; j <= r[i]; j++ {
			nj := nums[j]
			mi, ma = min(mi, nj), max(ma, nj)
		}
		if ma == mi {
			o[i] = true
		} else if (ma-mi)%(m-1) != 0 {
			o[i] = false
		} else {
			bs := make([]bool, m)
			j, d := l[i], (ma-mi)/(m-1)
			for ; j <= r[i]; j++ {
				d1 := nums[j] - mi
				if d1%d != 0 || bs[d1/d] {
					break
				}
				bs[d1/d] = true
			}
			o[i] = (j > r[i])
		}
	}
	return o
}
