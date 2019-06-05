func max(a, b string) string {
	if a < b {
		return b
	}
	return a
}

func reverseString(s string) string {
	bs := []byte(s)
	for i := 0; i < len(s)>>1; i++ {
		bs[i], bs[len(s)-i-1] = bs[len(s)-i-1], bs[i]
	}
	return string(bs)
}

func splitLoopedString(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	k, s, out := 0, "", "a"
	revs := make([]string, len(strs))
	for i, str := range strs {
		rev := reverseString(str)
		revs[i] = rev
		if str > rev {
			s += str
		} else {
			s += rev
		}
	}
	for i, str := range strs {
		rev, mid := revs[i], s[k+len(str):]+s[:k]
		for j := 0; j < len(str); j++ {
			if str[j] >= out[0] {
				out = max(out, str[j:]+mid+str[:j])
			}
			if rev[j] >= out[0] {
				out = max(out, rev[j:]+mid+rev[:j])
			}
		}
		k += len(str)
	}
	return out
}
