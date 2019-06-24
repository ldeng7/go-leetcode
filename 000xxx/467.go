func findSubstringInWraproundString(p string) int {
	if 0 == len(p) {
		return 0
	}
	cs, l := [26]int{}, 1
	cs[p[0]-'a'] = 1
	for i := 1; i < len(p); i++ {
		if p[i] == p[i-1]+1 || p[i-1]-p[i] == 25 {
			l++
		} else {
			l = 1
		}
		j := p[i] - 'a'
		if l > cs[j] {
			cs[j] = l
		}
	}
	o := 0
	for _, c := range cs {
		o += c
	}
	return o
}
