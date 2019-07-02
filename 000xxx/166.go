import (
	"bytes"
	"strconv"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func fractionToDecimal(numerator int, denominator int) string {
	n, d := abs(numerator), abs(denominator)
	q, r := n/d, n%d
	m := map[int]int{}
	o := strconv.Itoa(q)
	if (numerator >= 0) != (denominator >= 0) && (q > 0 || r > 0) {
		o = "-" + o
	}
	if r == 0 {
		return o
	}

	buf := &bytes.Buffer{}
	buf.WriteByte('.')
	i := 1
	for r != 0 {
		if j, ok := m[r]; ok {
			bs := buf.Bytes()
			bs = append(bs, 0)
			copy(bs[j+1:], bs[j:])
			bs[j] = '('
			bs = append(bs, ')')
			return o + string(bs)
		}
		buf.WriteString(strconv.Itoa((r * 10) / d))
		m[r], r, i = i, (r*10)%d, i+1
	}
	return o + buf.String()
}
