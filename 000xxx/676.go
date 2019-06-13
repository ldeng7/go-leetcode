type MagicDictionary struct {
	s map[string]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{map[string]bool{}}
}

func (this *MagicDictionary) BuildDict(dict []string) {
	for _, w := range dict {
		this.s[w] = true
	}
}

func (this *MagicDictionary) Search(word string) bool {
	bs := []byte(word)
	for i, c0 := range bs {
		for c := byte('a'); c <= 'z'; c++ {
			if c == c0 {
				continue
			}
			bs[i] = c
			if this.s[string(bs)] {
				return true
			}
		}
		bs[i] = c0
	}
	return false
}
