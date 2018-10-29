import "strings"

func findDuplicate(paths []string) [][]string {
	out := [][]string{}
	m := map[string][]string{}
	for _, path := range paths {
		es := strings.Split(path, " ")
		for _, e := range es[1:] {
			i := strings.LastIndexByte(e, '(')
			k := e[i+1 : len(e)-1]
			m[k] = append(m[k], es[0]+"/"+e[:i])
		}
	}
	for _, sl := range m {
		if len(sl) > 1 {
			out = append(out, sl)
		}
	}
	return out
}
