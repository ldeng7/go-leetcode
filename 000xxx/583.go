func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	t := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		t[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				t[i][j] = t[i-1][j-1] + 1
			} else if t[i-1][j] > t[i][j-1] {
				t[i][j] = t[i-1][j]
			} else {
				t[i][j] = t[i][j-1]
			}
		}
	}
	return l1 + l2 - t[l1][l2]*2
}
