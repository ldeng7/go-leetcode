import (
	"bytes"
	"strconv"
	"strings"
)

func s2f(s string) float64 {
	i := strings.IndexByte(s, '(')
	b := &bytes.Buffer{}
	if i >= 0 {
		b.WriteString(s[:i])
		t := s[i+1 : len(s)-1]
		for j := 0; j < 15; j++ {
			b.WriteString(t)
		}
	} else {
		b.WriteString(s)
	}
	f, _ := strconv.ParseFloat(b.String(), 64)
	return f
}

func isRationalEqual(S string, T string) bool {
	d := s2f(S) - s2f(T)
	if d < 0 {
		d = -d
	}
	return d < 0.00000001
}
