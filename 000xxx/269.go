func alienOrder(words []string) string {
	t := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		t[i] = make([]bool, 26)
	}
	v := make([]bool, 26)
	bs := []byte{}
	for _, w := range words {
		for _, c := range []byte(w) {
			t[c-'a'][c-'a'] = true
		}
	}
	for i := 0; i < len(words)-1; i++ {
		w0, w1 := words[i], words[i+1]
		l0, l1 := len(w0), len(w1)
		l := l0
		if l1 < l {
			l = l1
		}
		j := 0
		for ; j < l; j++ {
			c0, c1 := w0[j], w1[j]
			if c0 != c1 {
				t[c0-'a'][c1-'a'] = true
				break
			}
		}
		if j == l && l0 > l1 {
			return ""
		}
	}

	var cal func(idx byte) bool
	cal = func(idx byte) bool {
		if !t[idx][idx] {
			return true
		}
		v[idx] = true
		for i := byte(0); i < 26; i++ {
			if i == idx || !t[idx][i] {
				continue
			}
			if v[i] || (!cal(i)) {
				return false
			}
		}
		v[idx], t[idx][idx] = false, false
		bs = append(bs, 'a'+idx)
		return true
	}

	for i := byte(0); i < 26; i++ {
		if !cal(i) {
			return ""
		}
	}
	for i := 0; i < len(bs)>>1; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}
	return string(bs)
}
