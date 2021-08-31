func longestBeautifulSubstring(word string) int {
	o, l := 0, len(word)
	for i := 0; i < l; {
		if word[i] != 'a' {
			i++
			continue
		}
		j, c := i, byte('a')
		for ; j < l; j++ {
			c1 := word[j]
			if c1 == c {
				continue
			}
			var m byte
			switch c {
			case 'a':
				m = 'e'
			case 'e':
				m = 'i'
			case 'i':
				m = 'o'
			case 'o':
				m = 'u'
			case 'u':
				m = 'x'
			}
			if c1 == m {
				c = m
			} else {
				break
			}
		}
		if t := j - i; c == 'u' && t > o {
			o = t
		}
		i = j
	}
	return o
}
