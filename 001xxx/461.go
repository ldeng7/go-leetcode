func hasAllCodes(s string, k int) bool {
	if len(s) < k-1 {
		return false
	}

	var t uint32
	v := make([]bool, 1<<k)
	for i := 0; i < k-1; i++ {
		t <<= 1
		if s[i] == '1' {
			t |= 1
		}
	}
	for i := k - 1; i < len(s); i++ {
		t <<= 1
		if s[i] == '1' {
			t |= 1
		}
		v[t] = true
		t &^= 1 << (k - 1)
	}

	for _, b := range v {
		if !b {
			return false
		}
	}
	return true
}
