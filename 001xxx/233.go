import "sort"

func removeSubfolders(folder []string) []string {
	l := len(folder)
	o := make([]string, 1, l)
	sort.Strings(folder)
	p := folder[0]
	o[0] = p
	for i := 1; i < l; i++ {
		f := folder[i]
		lf, lp := len(f), len(p)
		if lf > lp && f[:lp] == p && f[lp] == '/' {
			continue
		}
		o = append(o, f)
		p = f
	}
	return o
}
