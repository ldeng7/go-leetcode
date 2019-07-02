func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	o, n := []byte{}, (numRows-1)<<1
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += n {
			o = append(o, s[j])
			k := j + n - i*2
			if i > 0 && i < numRows-1 && k < len(s) {
				o = append(o, s[k])
			}
		}
	}
	return string(o)
}
