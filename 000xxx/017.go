var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	ss := []string{}
	if 0 == len(digits) {
		return ss
	}
	bs := m[digits[0]]
	if 1 == len(digits) {
		for _, b := range bs {
			ss = append(ss, b)
		}
		return ss
	}
	subs := letterCombinations(digits[1:])
	for _, b := range bs {
		for _, s := range subs {
			ss = append(ss, b+s)
		}
	}
	return ss
}
