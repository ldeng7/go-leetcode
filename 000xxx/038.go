func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	out := []byte{'1'}
	for n > 1 {
		o := []byte{}
		for i := 0; i < len(out); i++ {
			var c byte = '1'
			for i < len(out)-1 && out[i] == out[i+1] {
				c, i = c+1, i+1
			}
			o = append(o, c, out[i])
		}
		n, out = n-1, o
	}
	return string(out)
}
