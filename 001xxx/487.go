import "fmt"

func getFolderNames(names []string) []string {
	o := make([]string, len(names))
	m := map[string]int{}
	for i, name := range names {
		if j, ok := m[name]; !ok {
			o[i], m[name] = name, 0
		} else {
			for {
				j++
				n := fmt.Sprintf("%s(%d)", name, j)
				if _, ok = m[n]; !ok {
					o[i], m[n] = n, 0
					break
				}
			}
			m[name] = j
		}
	}
	return o
}
