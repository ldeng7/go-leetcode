func rotatedDigits(N int) int {
	out := 0
	t := make([]int, N+1)
	i := 0
	for ; i <= N && i < 10; i++ {
		switch i {
		case 0, 1, 8:
			t[i] = 1
		case 2, 5, 6, 9:
			t[i], out = 2, out+1
		}
	}
	for ; i <= N; i++ {
		a, b := t[i/10], t[i%10]
		if a == 1 && b == 1 {
			t[i] = 1
		} else if a >= 1 && b >= 1 {
			t[i], out = 2, out+1
		}
	}
	return out
}
