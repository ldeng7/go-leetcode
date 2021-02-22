const M = 1000000007
const N = 13
const L = 10000

var m [L + N][N + 1]int

func init() {
	m[0][0] = 1
	for i := 1; i < L+N; i++ {
		m[i][0] = 1
		for j := 1; j <= i && j <= N; j++ {
			m[i][j] = (m[i-1][j-1] + m[i-1][j]) % M
		}
	}
}

func waysToFillArray(queries [][]int) []int {
	o := make([]int, len(queries))
	for i, q := range queries {
		n, k, s := q[0], q[1], 1
		for j := 2; j*j <= k; j++ {
			if k%j != 0 {
				continue
			}
			c := 0
			for ; k%j == 0; k, c = k/j, c+1 {
			}
			s = s * m[n+c-1][c] % M
		}
		if k != 1 {
			s = s * n % M
		}
		o[i] = s
	}
	return o
}
