import "math"

func lower(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	return 'a' + b - 'A'
}

func upper(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b
	}
	return 'A' + b - 'a'
}

func letterCasePermutation(s string) []string {
	bs := []byte(s)
	nc := 0
	for _, b := range bs {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			nc++
		}
	}
	if 0 == nc {
		return []string{s}
	}
	l := int(math.Pow(2.0, float64(nc)))
	ss := make([][]byte, l)
	for i := 0; i < l; i++ {
		ss[i] = make([]byte, len(s))
	}

	ic := 1 << uint(nc-1)
	for i, b := range bs {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			for j, s := range ss {
				if j&ic == 0 {
					s[i] = lower(b)
				} else {
					s[i] = upper(b)
				}
			}
			ic >>= 1
		} else {
			for _, s := range ss {
				s[i] = b
			}
		}
	}
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = string(s)
	}
	return out
}
