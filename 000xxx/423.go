func originalDigits(s string) string {
	bs := []byte{}
	cs, ns := [26]int{}, [10]int{}
	for _, b := range []byte(s) {
		cs[b-'a']++
	}
	ns[0] = cs['z'-'a']
	ns[2] = cs['w'-'a']
	ns[4] = cs['u'-'a']
	ns[6] = cs['x'-'a']
	ns[8] = cs['g'-'a']
	ns[1] = cs['o'-'a'] - ns[0] - ns[2] - ns[4]
	ns[3] = cs['h'-'a'] - ns[8]
	ns[5] = cs['f'-'a'] - ns[4]
	ns[7] = cs['s'-'a'] - ns[6]
	ns[9] = cs['i'-'a'] - ns[6] - ns[8] - ns[5]
	for i, n := range ns {
		for j := 0; j < n; j++ {
			bs = append(bs, byte(i)+'0')
		}
	}
	return string(bs)
}
