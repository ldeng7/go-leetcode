func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	out := []byte{}
	n := (numRows - 1) * 2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += n {
			out = append(out, s[j])
			k := j + n - i*2
			if i > 0 && i < numRows-1 && k < len(s) {
				out = append(out, s[k])
			}
		}
	}
	return string(out)
}