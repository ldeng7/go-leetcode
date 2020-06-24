func numberOfSubstrings(s string) int {
	o, l, t := 0, len(s), 0
	ar := [4]int{}
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		ar[c]++
		for ar[0] != 0 && ar[1] != 0 && ar[2] != 0 {
			ar[s[t]-'a']--
			t++
		}
		o += i - t + 1
	}
	return (l*(l+1))>>1 - o
}
