var m map[byte]byte = map[byte]byte{
	0: 0, 1: 1, 6: 9, 8: 8, 9: 6,
}

func confusingNumber(N int) bool {
	if 0 == N {
		return false
	}
	ds := []byte{}
	for N > 0 {
		ds = append(ds, byte(N%10))
		N /= 10
	}
	for _, d := range ds {
		if _, ok := m[d]; !ok {
			return false
		}
	}
	for i := 0; i <= len(ds)>>1; i++ {
		if m[ds[i]] != ds[len(ds)-i-1] {
			return true
		}
	}
	return false
}
