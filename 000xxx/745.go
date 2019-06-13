type WordFilter struct {
	m map[string]int
}

func Constructor(words []string) WordFilter {
	m := map[string]int{}
	for k, word := range words {
		l := len(word)
		for i := 0; i <= l; i++ {
			for j := 0; j <= l; j++ {
				m[word[:i]+";"+word[j:]] = k
			}
		}
	}
	return WordFilter{m}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if k, ok := this.m[prefix+";"+suffix]; ok {
		return k
	}
	return -1
}
