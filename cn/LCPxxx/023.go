func isMagic(target []int) bool {
	l := len(target)
	ar := make([]int, 0, l)
	for i := 2; i <= l; i += 2 {
		ar = append(ar, i)
	}
	for i := 1; i <= l; i += 2 {
		ar = append(ar, i)
	}
	k := 0
	for ; k < l; k++ {
		if ar[k] != target[k] {
			break
		}
	}
	if 0 == k {
		return false
	}

	res := make([]int, 0, l)
	for len(ar) > 0 {
		d := k
		if d > len(ar) {
			d = len(ar)
		}
		res = append(res, ar[:d]...)
		ar = ar[d:]
		t := make([]int, 0, len(ar))
		for i := 1; i < len(ar); i += 2 {
			t = append(t, ar[i])
		}
		for i := 0; i < len(ar); i += 2 {
			t = append(t, ar[i])
		}
		ar = t
	}
	for i, tar := range target {
		if res[i] != tar {
			return false
		}
	}
	return true
}
