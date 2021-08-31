func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	o := make([]byte, 0, l1+l2)
	i, j := 0, 0
	for ; i < l1 && j < l2; i, j = i+1, j+1 {
		o = append(o, word1[i], word2[j])
	}
	for ; i < l1; i++ {
		o = append(o, word1[i])
	}
	for ; i < l2; i++ {
		o = append(o, word2[i])
	}
	return string(o)
}
