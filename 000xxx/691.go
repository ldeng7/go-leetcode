import "math"

func minStickers(stickers []string, target string) int {
	l := uint(len(target))
	m := 1 << l
	t := make([]int, m)
	for i := 1; i < m; i++ {
		t[i] = math.MaxInt64
	}
	for i := 0; i < m; i++ {
		if t[i] == math.MaxInt64 {
			continue
		}
		for _, s := range stickers {
			ic := i
			for j := 0; j < len(s); j++ {
				c := s[j]
				for k := uint(0); k < l; k++ {
					if target[k] == c && (ic>>k)&1 == 0 {
						ic |= 1 << k
						break
					}
				}
			}
			if t[i]+1 < t[ic] {
				t[ic] = t[i] + 1
			}
		}
	}
	if t[m-1] == math.MaxInt64 {
		return -1
	}
	return t[m-1]
}
