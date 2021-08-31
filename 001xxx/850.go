func getMinSwaps(num string, k int) int {
	bs := []byte(num)
	l := len(bs)
	for i := 0; i < k; i++ {
		j := l - 2
		for ; j >= 0 && bs[j] >= bs[j+1]; j-- {
		}
		if j >= 0 {
			k := l - 1
			for ; k >= 0 && bs[k] <= bs[j]; k-- {
			}
			bs[j], bs[k] = bs[k], bs[j]
		}
		for k, m := j+1, l-1; k < m; k, m = k+1, m-1 {
			bs[k], bs[m] = bs[m], bs[k]
		}
	}

	o := 0
	for i := 0; i < l; i++ {
		if num[i] == bs[i] {
			continue
		}
		j := i + 1
		for ; num[i] != bs[j]; j++ {
		}
		for ; j != i; j-- {
			bs[j], bs[j-1] = bs[j-1], bs[j]
			o++
		}
	}
	return o
}
