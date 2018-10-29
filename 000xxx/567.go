func checkInclusion(s1 string, s2 string) bool {
	cnt, l := len(s1), 0
	ar := make([]int, 128)
	for _, b := range []byte(s1) {
		ar[b]++
	}
	for r := 0; r < len(s2); r++ {
		if ar[s2[r]] > 0 {
			cnt--
		}
		ar[s2[r]]--
		for 0 == cnt {
			if r-l+1 == len(s1) {
				return true
			}
			ar[s2[l]]++
			if ar[s2[l]] > 0 {
				cnt++
			}
			l++
		}
	}
	return false
}
