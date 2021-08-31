func largestMerge(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	o := make([]byte, 0, m+n)
	i1, i2 := 0, 0
	for i1 < m || i2 < n {
		if word1[i1:] > word2[i2:] {
			o = append(o, word1[i1])
			i1++
		} else {
			o = append(o, word2[i2])
			i2++
		}
	}
	return string(o)
}
