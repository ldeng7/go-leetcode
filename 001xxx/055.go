func shortestWay(source string, target string) int {
	ls, lt := len(source), len(target)
	m := [26]bool{}
	for i := 0; i < ls; i++ {
		m[source[i]-'a'] = true
	}

	o := 1
	for i, j := 0, 0; i < lt; {
		ct := target[i]
		if idx := ct - 'a'; !m[idx] {
			return -1
		}

		if j == ls {
			j, o = 0, o+1
		}
		for j < ls && ct != source[j] {
			j++
		}
		if j != ls {
			i, j = i+1, j+1
		}
	}
	return o
}
