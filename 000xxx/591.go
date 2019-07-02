import "strings"

func isValid(code string) bool {
	st := []string{}
	l := len(code)
	for i := 0; i < l; i++ {
		if i > 0 && 0 == len(st) {
			return false
		}
		if l-i >= 9 && code[i:i+9] == "<![CDATA[" {
			j := i + 9
			i = strings.Index(code[j:], "]]>") + j
			if i < j {
				return false
			}
			i += 2
		} else if l-i >= 2 && code[i:i+2] == "</" {
			j := i + 2
			i = strings.IndexByte(code[j:], '>') + j
			if i < j || 0 == len(st) || st[len(st)-1] != code[j:i] {
				return false
			}
			st = st[:len(st)-1]
		} else if code[i] == '<' {
			j := i + 1
			i = strings.IndexByte(code[j:], '>') + j
			if i <= j || i-j > 9 {
				return false
			}
			for k := j; k < i; k++ {
				if code[k] < 'A' || code[k] > 'Z' {
					return false
				}
			}
			st = append(st, code[j:i])
		}
	}
	return 0 == len(st)
}
