import "bytes"

func areAlmostEqual(s1 string, s2 string) bool {
	o, k, l := true, -1, len(s1)
	bs1, bs2 := []byte(s1), []byte(s2)
	for i := 0; i < l; i++ {
		c1, c2 := bs1[i], bs2[i]
		if c1 != c2 && o {
			o, k = false, i
			continue
		}
		if c1 != c2 && !o {
			bs1[k], bs1[i] = c1, bs1[k]
			return bytes.Compare(bs1, bs2) == 0
		}
	}
	return o
}
