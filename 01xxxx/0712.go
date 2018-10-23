func minimumDeleteSum(s1 string, s2 string) int {
	t := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		t[i] = make([]int, len(s2)+1)
	}
	for i := 1; i <= len(s2); i++ {
		t[0][i] = t[0][i-1] + int(s2[i-1])
	}
	for i := 1; i <= len(s1); i++ {
		t[i][0] = t[i-1][0] + int(s1[i-1])
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				t[i][j] = t[i-1][j-1]
			} else {
				a1, a2 := t[i-1][j]+int(s1[i-1]), t[i][j-1]+int(s2[j-1])
				if a1 < a2 {
					t[i][j] = a1
				} else {
					t[i][j] = a2
				}
			}
		}
	}
	return t[len(s1)][len(s2)]
}
