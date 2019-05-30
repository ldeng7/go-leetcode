import "strings"

func parseTernary(expression string) string {
	out := expression
	for len(out) > 1 {
		i := strings.LastIndexByte(out, '?')
		if out[i-1] == 'T' {
			out = out[:i-1] + out[i+1:i+2] + out[i+4:]
		} else {
			out = out[:i-1] + out[i+3:]
		}
	}
	return out
}
