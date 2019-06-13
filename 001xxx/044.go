import "strings"

const prime = 16777619

func longestDupSubstring(S string) string {
	n := len(S)
	l, r := 0, n-1
	var out string
	for l <= r {
		m, t := l+(r-l)>>1, map[uint32]int{}
		var h, p, sq uint32 = 0, 1, prime
		for i := 0; i < m; i++ {
			h = h*prime + uint32(S[i])
		}
		for i := m; i > 0; i >>= 1 {
			if i&1 != 0 {
				p *= sq
			}
			sq *= sq
		}
		t[h] = 0

		i := m
		for i < n {
			h = (h * prime) + uint32(S[i]) - p*uint32(S[i-m])
			i++
			if j, ok := t[h]; ok {
				sub := S[i-m : i]
				if sub == S[j:j+m] || strings.Index(S[:i-1], sub) >= 0 {
					out, l = sub, m+1
					break
				}
			}
			t[h] = i - m
		}
		if i == n {
			r = m - 1
		}
	}
	return out
}
