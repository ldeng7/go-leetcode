var m map[byte]byte = map[byte]byte{
	'0': '0', '1': '1', '6': '9', '8': '8', '9': '6',
}

func isStrobogrammatic(num string) bool {
	for i := 0; i <= len(num)>>1; i++ {
		if m[num[i]] != num[len(num)-i-1] {
			return false
		}
	}
	return true
}
