import "strings"

func simplifyPath(path string) string {
	st := []string{}
	parts := strings.Split(path, "/")
	for _, p := range parts {
		if 0 != len(st) && p == ".." {
			st = st[:len(st)-1]
		} else if 0 != len(p) && p != "." && p != ".." {
			st = append(st, p)
		}
	}
	return "/" + strings.Join(st, "/")
}
