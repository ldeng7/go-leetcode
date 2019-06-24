func wordSubsets(A []string, B []string) []string {
	o := []string{}
	cs := [26]int{}
	for _, s := range B {
		cs1 := [26]int{}
		for i := 0; i < len(s); i++ {
			cs1[s[i]-'a']++
		}
		for i := 0; i < 26; i++ {
			if cs1[i] > cs[i] {
				cs[i] = cs1[i]
			}
		}
	}
	for _, s := range A {
		cs1 := [26]int{}
		for i := 0; i < len(s); i++ {
			cs1[s[i]-'a']++
		}
		i := 0
		for ; i < 26; i++ {
			if cs1[i] < cs[i] {
				break
			}
		}
		if i == 26 {
			o = append(o, s)
		}
	}
	return o
}
