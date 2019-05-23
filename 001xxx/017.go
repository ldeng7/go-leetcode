func reverseBytes(bs []byte) []byte {
	bsr := make([]byte, len(bs))
	for i, b := range []byte(bs) {
		bsr[len(bs)-i-1] = b
	}
	return bsr
}

func baseNeg2(N int) string {
	bs := []byte{}
	for 0 != N {
		bs = append(bs, '0'+byte(N&1))
		N = -(N >> 1)
	}
	if 0 == len(bs) {
		return "0"
	}
	return string(reverseBytes(bs))
}
