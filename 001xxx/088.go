var dm = map[byte]byte{
	0: 0, 1: 1, 6: 9, 8: 8, 9: 6,
}

func confusingNumberII(N int) int {
	o := 0
	ds := make([]byte, 0, 10)
	for ; N != 0; N /= 10 {
		ds = append(ds, byte(N%10))
	}
	l := len(ds)
	for i := 0; i < l>>1; i++ {
		ds[i], ds[l-i-1] = ds[l-i-1], ds[i]
	}
	t := []byte{}

	var cal func(int, bool)
	cal = func(i int, all bool) {
		if i == l {
			j := 0
			for ; j < len(t) && t[j] == 0; j++ {
			}
			t1, lt := t[j:], len(t)-j
			for k := 0; k < lt; k++ {
				if t1[k] != dm[t1[lt-k-1]] {
					o++
					break
				}
			}
			return
		}

		var je byte = 9
		if !all {
			je = ds[i]
		}
		for j := byte(0); j <= je; j++ {
			if j >= 2 && j <= 7 && j != 6 {
				continue
			}
			t = append(t, j)
			cal(i+1, all || j != je)
			t = t[:len(t)-1]
		}
	}

	cal(0, false)
	return o
}
