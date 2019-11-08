func canConvert(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	l1, l2 := len(str1), len(str2)
	c, m := 0, [26]bool{}
	for i := 0; i < l2; i++ {
		j := str2[i] - 'a'
		if !m[j] {
			m[j], c = true, c+1
		}
	}
	if c == 26 {
		return false
	}
	ar := [26]int{}
	for i := 0; i < l1; i++ {
		j := str1[i] - 'a'
		if ar[j] != 0 && str2[ar[j]-1] != str2[i] {
			return false
		}
		ar[j] = i + 1
	}
	return true
}
