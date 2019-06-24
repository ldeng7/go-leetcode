func lexicalOrder(n int) []int {
	o := make([]int, n)
	j := 1
	for i := 0; i < n; i++ {
		o[i] = j
		if j*10 <= n {
			j *= 10
		} else {
			if j >= n {
				j /= 10
			}
			j += 1
			for j%10 == 0 {
				j /= 10
			}
		}
	}
	return o
}
