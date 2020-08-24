func winnerSquareGame(n int) bool {
	ar := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= 316; j++ {
			if k := i - j*j; k < 0 {
				break
			} else if !ar[k] {
				ar[i] = true
				break
			}
		}
	}
	return ar[n]
}
