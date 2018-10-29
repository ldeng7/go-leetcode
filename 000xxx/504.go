func convertToBase7(num int) string {
	m := false
	if num < 0 {
		num = -num
		m = true
	}
	bs := []byte{}
	for num > 0 {
		bs = append(bs, '0'+byte(num%7))
		num /= 7
	}
	if 0 == len(bs) {
		bs = []byte{'0'}
	} else if m {
		bs = append(bs, '-')
	}
	bsr := make([]byte, len(bs))
	j := 0
	for i := len(bs) - 1; i >= 0; i-- {
		bsr[j] = bs[i]
		j++
	}
	return string(bsr)
}
