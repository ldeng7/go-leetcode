func distance(a, b int) int {
	if a <= b {
		return b - a
	}
	return a - b
}

type WordDistance struct {
	m map[string][]int
}

func Constructor(words []string) WordDistance {
	wd := &WordDistance{map[string][]int{}}
	for i, w := range words {
		if ar, ok := wd.m[w]; ok {
			wd.m[w] = append(ar, i)
		} else {
			wd.m[w] = []int{i}
		}
	}
	return *wd
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	indices := []int{}
	classes := []bool{}
	ar1, ar2 := this.m[word1], this.m[word2]
	j1, j2 := 0, 0
	for j1 < len(ar1) && j2 < len(ar2) {
		i1, i2 := ar1[j1], ar2[j2]
		if i1 < i2 {
			indices = append(indices, i1)
			classes = append(classes, true)
			j1++
		} else {
			indices = append(indices, i2)
			classes = append(classes, false)
			j2++
		}
	}
	if j1 < len(ar1) {
		indices = append(indices, ar1[j1])
		classes = append(classes, true)
	} else {
		indices = append(indices, ar2[j2])
		classes = append(classes, false)
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
