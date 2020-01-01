import "strconv"

func toHexspeak(num string) string {
	i, _ := strconv.ParseInt(num, 10, 64)
	s := strconv.FormatInt(i, 16)
	bs := []byte(s)
	for i, b := range bs {
		switch b {
		case '0':
			bs[i] = 'O'
		case '1':
			bs[i] = 'I'
		case 'a', 'b', 'c', 'd', 'e', 'f':
			bs[i] = b - 32
		default:
			return "ERROR"
		}
	}
	return string(bs)
}
