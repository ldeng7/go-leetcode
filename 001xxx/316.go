func distinctEchoSubstrings(text string) int {
	l := len(text)
	m := l - 1
	set := map[string]bool{}
	for i := 0; i <= m; i++ {
		n := l - i
		ar := make([]int, n)
		for j, k := 1, 0; j < n; j++ {
			c := text[i+j]
			for k > 0 && c != text[i+k] {
				k = ar[k-1]
			}
			if c == text[i+k] {
				k++
			}
			ar[j] = k
			if k == j {
				if k&1 == 1 {
					set[text[i:i+((j+1)>>1)]] = true
				}
			} else if t := j + 1 - k; k != 0 && (j+1)%t == 0 {
				if ((j+1)/t)&1 == 0 {
					set[text[i:i+((j+1)>>1)]] = true
				}
			}
		}
		t := 2000
		if t1 := n - ar[n-1]; ar[n-1] != 0 && n%t1 == 0 {
			t = t1
		}
		if t < m {
			m = t
		}
	}
	return len(set)
}
