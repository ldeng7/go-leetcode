import "bytes"

func alphabetBoardPath(target string) string {
	y, x := 0, 0
	buf := &bytes.Buffer{}
	for i := 0; i < len(target); i++ {
		if i > 0 && target[i] == target[i-1] {
			buf.WriteByte('!')
			continue
		}
		b := int(target[i] - 'a')
		dy, dx := b/5, b%5
		if dx < x {
			for j := 0; j < x-dx; j++ {
				buf.WriteByte('L')
			}
		}
		if dy < y {
			for j := 0; j < y-dy; j++ {
				buf.WriteByte('U')
			}
		}
		if dx > x {
			for j := 0; j < dx-x; j++ {
				buf.WriteByte('R')
			}
		}
		if dy > y {
			for j := 0; j < dy-y; j++ {
				buf.WriteByte('D')
			}
		}
		buf.WriteByte('!')
		y, x = dy, dx
	}
	return buf.String()
}
