func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	m := map[string]string{}
	for _, rs := range regions {
		l, r0 := len(rs), rs[0]
		for i := 1; i < l; i++ {
			m[rs[i]] = r0
		}
	}
	m1, m2 := map[string]bool{}, map[string]bool{}
	m1[region1], m2[region2] = true, true
	for {
		m1[m[region1]], m2[m[region2]] = true, true
		if o := m[region2]; m1[o] {
			return o
		}
		if o := m[region1]; m2[o] {
			return o
		}
		region1, region2 = m[region1], m[region2]
	}
	return ""
}
