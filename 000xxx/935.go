var m = [10][]int{
	{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 9, 3}, {}, {1, 7, 0}, {2, 6}, {1, 3}, {2, 4},
}

func knightDialer(N int) int {
	o := 0
	t := make([][10]int, N)
	for j := 0; j <= 9; j++ {
		t[0][j] = 1
	}
	for i := 1; i < N; i++ {
		for j := 0; j <= 9; j++ {
			for _, k := range m[j] {
				t[i][j] = (t[i][j] + t[i-1][k]) % 1000000007
			}
		}
	}
	for i := 0; i <= 9; i++ {
		o = (o + t[N-1][i]) % 1000000007
	}
	return o
}
