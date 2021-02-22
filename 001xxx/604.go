import "sort"

func alertNames(keyName []string, keyTime []string) []string {
	m := map[string][]int{}
	for i, n := range keyName {
		t := keyTime[i]
		m[n] = append(m[n], (int(t[0]-'0')*10+int(t[1]-'0'))*60+int(t[3]-'0')*10+int(t[4]-'0'))
	}
	o := make([]string, 0, 16)
	for n, ar := range m {
		l := len(ar)
		if l <= 2 {
			continue
		}
		sort.Ints(ar)
		for i := 2; i < l; i++ {
			if ar[i]-ar[i-2] <= 60 {
				o = append(o, n)
				break
			}
		}
	}
	sort.Strings(o)
	return o
}
