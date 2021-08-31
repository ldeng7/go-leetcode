import "math/bits"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func countPairs(nums []int, low int, high int) int {
	o, n, m, l, r := 0, len(nums), 0, low, high+1
	for _, a := range nums {
		m = max(m, 16-bits.LeadingZeros16(uint16(a)))
	}
	m = 1 << m
	if r >= m {
		o, r = n*n, 0
		if l >= m {
			o, l = 0, 0
		}
	}

	ar, ar1 := make([]int, m), make([]int, m>>1)
	for i := 0; i < m; i++ {
		ar[i] = 0
	}
	for _, a := range nums {
		ar[a]++
	}
	for ; l != r; l, r, m = l>>1, r>>1, m>>1 {
		for i, m1 := 0, m>>1; i < m1; i++ {
			ar1[i] = 0
		}
		if (l&1 != 0) && (r&1 != 0) {
			for i := 0; i < m; i++ {
				ar1[i>>1] += ar[i]
				o += ar[i] * (ar[(r-1)^i] - ar[(l-1)^i])
			}
		} else if l&1 != 0 {
			for i := 0; i < m; i++ {
				ar1[i>>1] += ar[i]
				o -= ar[i] * ar[(l-1)^i]
			}
		} else if r&1 != 0 {
			for i := 0; i < m; i++ {
				ar1[i>>1] += ar[i]
				o += ar[i] * ar[(r-1)^i]
			}
		} else {
			for i := 0; i < m; i++ {
				ar1[i>>1] += ar[i]
			}
		}
		ar, ar1 = ar1, ar
	}
	return o / 2
}
