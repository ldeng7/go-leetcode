func compareVersion(version1 string, version2 string) int {
	l1, l2 := len(version1), len(version2)
	a, b := 0, 0
	for i, j := 0, 0; i < l1 || j < l2; i, j = i+1, j+1 {
		for i < l1 && version1[i] != '.' {
			a = a*10 + int(version1[i]-'0')
			i++
		}
		for j < l2 && version2[j] != '.' {
			b = b*10 + int(version2[j]-'0')
			j++
		}
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		a, b = 0, 0
	}
	return 0
}
