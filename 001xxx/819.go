func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func countDifferentSubsequenceGCDs(nums []int) int {
	ma := 0
	m := make([]bool, 200001)
	for _, n := range nums {
		ma = max(ma, n)
		m[n] = true
	}
	o := 0
	for i := 1; i <= ma; i++ {
		t := 0
		for j := i; j <= ma; j += i {
			if !m[j] {
				continue
			}
			t = gcd(t, j)
			if t == i {
				break
			}
		}
		if t == i {
			o++
		}
	}
	return o
}
