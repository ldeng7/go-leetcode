func canConvertString(s string, t string, k int) bool {
	l := len(s)
	if l != len(t) {
		return false
	}
	ar := [26]int{}
	for i := 0; i < l; i++ {
		d := int(t[i]) - int(s[i])
		if d < 0 {
			d += 26
		}
		if d != 0 && (d+ar[d]*26) > k {
			return false
		}
		ar[d]++
	}
	return true
}
