func maxCoins(nums []int) int {
	l := len(nums)
	ar := make([]int, l+2)
	ar[0], ar[l+1] = 1, 1
	copy(ar[1:], nums)
	t := make([][]int, l+2)
	for i := 0; i < l+2; i++ {
		t[i] = make([]int, l+2)
	}
	for i := 1; i <= l; i++ {
		for j := 1; j <= l-i+1; j++ {
			h := j + i - 1
			for k := j; k <= h; k++ {
				v := ar[j-1]*ar[k]*ar[h+1] + t[j][k-1] + t[k+1][h]
				if v > t[j][h] {
					t[j][h] = v
				}
			}
		}
	}
	return t[1][l]
}
