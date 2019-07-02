import "strings"

func devowel(s string) string {
	bs := []byte(s)
	for i, c := range bs {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			bs[i] = '*'
		}
	}
	return string(bs)
}

func spellchecker(wordlist []string, queries []string) []string {
	matches, caps, vowels := map[string]bool{}, map[string]string{}, map[string]string{}
	for _, w := range wordlist {
		matches[w] = true
		wl := strings.ToLower(w)
		if _, ok := caps[wl]; !ok {
			caps[wl] = w
		}
		wd := devowel(wl)
		if _, ok := vowels[wd]; !ok {
			vowels[wd] = w
		}
	}
	o := make([]string, len(queries))
	for i, q := range queries {
		if matches[q] {
			o[i] = q
			continue
		}
		ql := strings.ToLower(q)
		if w, ok := caps[ql]; ok {
			o[i] = w
		} else if w, ok = vowels[devowel(ql)]; ok {
			o[i] = w
		} else {
			o[i] = ""
		}
	}
	return o
}
