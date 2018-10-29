func removeKdigits(num string, k int) string {
	out := []byte{}
	r := len(num) - k
	for _, b := range []byte(num) {
		for k > 0 && len(out) > 0 && out[len(out)-1] > b {
			out = out[:len(out)-1]
			k--
		}
		out = append(out, b)
	}
	i := 0
	for ; i < len(out); i++ {
		if out[i] != '0' {
			break
		}
	}
	if i >= r {
		return "0"
	}
	return string(out[i:r])
}
