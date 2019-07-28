var m = map[byte]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}

func romanToInt(s string) int {
	o, l := 0, len(s)
	for i := 0; i < l; i++ {
		v := m[s[i]]
		if i == l-1 || v >= m[s[i+1]] {
			o += v
		} else {
			o -= v
		}
	}
	return o
}
