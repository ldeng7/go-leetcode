func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxLength(arr []string) int {
	o, l := 0, len(arr)
	t := make([]uint32, l)
	for i, s := range arr {
		for j := 0; j < len(s); j++ {
			var m uint32 = 1 << (s[j] - 'a')
			if t[i]&m == 0 {
				t[i] |= m
			} else {
				t[i], arr[i] = 0xfffffff, ""
				break
			}
		}
	}

	var cal func(int, int, uint32) int
	cal = func(i, n int, v uint32) int {
		if i == l {
			return n
		}
		if v&t[i] != 0 {
			return cal(i+1, n, v)
		}
		return max(cal(i+1, n+len(arr[i]), v|t[i]), cal(i+1, n, v))
	}
	for i, s := range arr {
		if o1 := cal(i+1, len(s), t[i]); o1 > o {
			o = o1
		}
	}
	return o
}
