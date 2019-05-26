func distance(a, b int) int {
	if a <= b {
		return b - a
	}
	return a - b
}

func shortestWordDistance(words []string, word1 string, word2 string) int {
	d := -1
	if word1 != word2 {
		indices := []int{}
		classes := []bool{}
		for i, w := range words {
			if w == word1 || w == word2 {
				indices = append(indices, i)
				classes = append(classes, w == word1)
			}
		}

		for i := 1; i < len(indices); i++ {
			if classes[i] != classes[i-1] {
				d1 := distance(indices[i], indices[i-1])
				if d1 < d || d == -1 {
					d = d1
				}
			}
		}
	} else {
		j := -1
		for i, w := range words {
			if w == word1 {
				if j != -1 {
					d1 := distance(i, j)
					if d1 < d || d == -1 {
						d = d1
					}
				}
				j = i
			}
		}
	}
	return d
}
