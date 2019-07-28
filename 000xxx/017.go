var m = [9][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"},
	{"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	o := []string{}
	if 0 == len(digits) {
		return o
	}
	bs := m[digits[0]-'2']
	if 1 == len(digits) {
		for _, b := range bs {
			o = append(o, b)
		}
		return o
	}
	subs := letterCombinations(digits[1:])
	for _, b := range bs {
		for _, s := range subs {
			o = append(o, b+s)
		}
	}
	return o
}
