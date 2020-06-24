func findTheLongestSubstring(s string) int {
	h, t := [32]int{}, [32]int{}
	h[0], t[0] = -1, -1
	for i := 1; i < 32; i++ {
		h[i], t[i] = -2, -2
	}
	l, v := len(s), 0
	for i := 0; i < l; i++ {
		switch s[i] {
		case 'a':
			v ^= 1
		case 'e':
			v ^= 2
		case 'i':
			v ^= 4
		case 'o':
			v ^= 8
		case 'u':
			v ^= 16
		}
		if -2 == h[v] {
			h[v] = i
		}
		t[v] = i
	}
	o := 0
	for i := 0; i < 32; i++ {
		if h[i] != -2 && t[i] != -2 {
			if o1 := t[i] - h[i]; o1 > o {
				o = o1
			}
		}
	}
	return o
}
