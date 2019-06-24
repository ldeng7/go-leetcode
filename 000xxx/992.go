func subarraysWithKDistinct(A []int, K int) int {
	o, l, d, j := 0, len(A), 0, 0
	t := make([]int, l+1)
	for i := 0; i < l; i++ {
		if i > 0 {
			a := A[i-1]
			t[a]--
			if t[a] == 0 {
				d--
			}
		}
		for j >= i && d >= K {
			j--
			a := A[j]
			t[a]--
			if t[a] == 0 {
				d--
			}
		}
		for ; j < l; j++ {
			if d > K {
				break
			}
			a := A[j]
			if t[a] == 0 {
				d++
			}
			t[a]++
			if d == K {
				o++
			}
		}
	}
	return o
}
