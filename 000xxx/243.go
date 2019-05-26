func distance(a, b int) int {
	if a <= b {
		return b - a
	}
	return a - b
}

func shortestDistance(words []string, word1 string, word2 string) int {
	indices := []int{}
	classes := []bool{}
	for i, w := range words {
		if w == word1 || w == word2 {
			indices = append(indices, i)
			classes = append(classes, w == word1)
		}
	}
	d := -1
	for i := 1; i < len(indices); i++ {
		if classes[i] != classes[i-1] {
			d1 := distance(indices[i], indices[i-1])
			if d1 < d || d == -1 {
				d = d1
			}
		}
	}
	return d
}
