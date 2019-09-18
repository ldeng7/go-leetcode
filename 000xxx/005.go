func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	rs := make([]int, l<<1+1)
	c, r, m := -1, -1, 0
	var o string
	for i := 0; i < l<<1+1; i++ {
		if i < r && c<<1-i >= 0 {
			rs[i] = min(r-i+1, rs[c<<1-i])
		} else {
			rs[i] = 1
		}
		for j := rs[i]; j <= i && i+j <= l<<1; j++ {
			if (i-j)&1 == 0 || s[(i-j)>>1] == s[(i+j)>>1] {
				rs[i]++
				c, r = i, i+j
			} else {
				break
			}
		}
		if rs[i] > m {
			m = rs[i]
			o = s[(i-rs[i]+1)>>1 : (i+rs[i]-1)>>1]
		}
	}
	return o
}
