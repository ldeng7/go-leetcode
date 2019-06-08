import "bytes"

func pushDominoes(dominoes string) string {
	buf := bytes.Buffer{}
	dominoes = "L" + dominoes + "R"
	for i, j := 0, 1; j < len(dominoes); j++ {
		ci, cj := dominoes[i], dominoes[j]
		if cj == '.' {
			continue
		}
		if i > 0 {
			buf.WriteByte(ci)
		}
		m := j - i - 1
		if ci == cj {
			for k := 0; k < m; k++ {
				buf.WriteByte(ci)
			}
		} else if ci == 'L' && cj == 'R' {
			for k := 0; k < m; k++ {
				buf.WriteByte('.')
			}
		} else {
			for k := 0; k < m>>1; k++ {
				buf.WriteByte('R')
			}
			if m&1 == 1 {
				buf.WriteByte('.')
			}
			for k := 0; k < m>>1; k++ {
				buf.WriteByte('L')
			}
		}
		i = j
	}
	return buf.String()
}
