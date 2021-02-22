func maximumBinaryString(binary string) string {
	l := len(binary)
	i, c := 0, 0
	for ; i < l && binary[i] == '1'; i++ {
	}
	for j := i; j < l; j++ {
		if binary[j] == '0' {
			c++
		}
	}
	if c < 2 {
		return binary
	}
	bs := []byte(binary)
	for ; c > 1; i, c = i+1, c-1 {
		bs[i] = '1'
	}
	bs[i] = '0'
	for i++; i < l; i++ {
		bs[i] = '1'
	}
	return string(bs)
}
