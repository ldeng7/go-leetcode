func removeComments(source []string) []string {
	out := []string{}
	b := true
	lo := []byte{}
	for _, l := range source {
		bs := []byte(l)
		i := 0
		for ; i < len(bs)-1; i++ {
			if b {
				if bs[i] == '/' && bs[i+1] == '*' {
					b = false
					i++
				} else if bs[i] == '/' && bs[i+1] == '/' {
					break
				} else {
					lo = append(lo, bs[i])
				}
			} else if bs[i] == '*' && bs[i+1] == '/' {
				b = true
				i++
			}
		}
		if b && i == len(bs)-1 {
			lo = append(lo, bs[len(bs)-1])
		}
		if len(lo) > 0 && b {
			out = append(out, string(lo))
			lo = []byte{}
		}
	}
	return out
}
