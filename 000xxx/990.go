func equationsPossible(equations []string) bool {
	t := [26][]byte{}
	for _, e := range equations {
		if e[1] != '=' {
			continue
		}
		i, j := e[0]-'a', e[3]-'a'
		t[i], t[j] = append(t[i], j), append(t[j], i)
	}
	cs := [26]byte{}
	for i, j := byte(0), byte(0); i < 26; i++ {
		if cs[i] != 0 {
			continue
		}
		j++
		st := []byte{i}
		for 0 != len(st) {
			e := len(st) - 1
			k := st[e]
			st = st[:e]
			for _, h := range t[k] {
				if cs[h] == 0 {
					cs[h], st = j, append(st, h)
				}
			}
		}
	}

	for _, e := range equations {
		if e[1] != '!' {
			continue
		}
		i, j := e[0]-'a', e[3]-'a'
		if i == j || cs[i] != 0 && cs[i] == cs[j] {
			return false
		}
	}

	return true
}
