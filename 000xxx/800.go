import "strconv"

func similarRGB(color string) string {
	bs := []byte(color)
	for i := 1; i < len(bs); i += 2 {
		n, _ := strconv.ParseUint(string(bs[i:i+2]), 16, 8)
		n1 := n / 17
		if n%17 > 8 {
			n1++
		}
		var b byte = byte(n1) + '0'
		if n1 > 9 {
			b = byte(n1) - 10 + 'a'
		}
		bs[i], bs[i+1] = b, b
	}
	return string(bs)
}
