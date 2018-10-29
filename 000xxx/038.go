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
				c++
				i++
			}
			o = append(o, c)
			o = append(o, out[i])
		}
		n--
		out = o
	}
	return string(out)
}