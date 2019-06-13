func cmp(sa, sb string) int {
	c := 0
	for i := 0; i < 6; i++ {
		if sa[i] == sb[i] {
			c++
		}
	}
	return c
}

func findSecretWord(wordlist []string, master *Master) {
	l := len(wordlist)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		t[i][i] = 6
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			c := cmp(wordlist[i], wordlist[j])
			t[i][j], t[j][i] = c, c
		}
	}

	m, j := l, 0
	for i := 0; i < l; i++ {
		cs, mc := [7]int{}, 0
		for _, j := range t[i] {
			cs[j]++
			if cs[j] > mc {
				mc = cs[j]
			}
		}
		if mc < m {
			m, j = mc, i
		}
	}

	wg := wordlist[j]
	res := master.Guess(wg)
	if 6 != res {
		sub := []string{}
		for _, w := range wordlist {
			if w != wg && cmp(w, wg) == res {
				sub = append(sub, w)
			}
		}
		findSecretWord(sub, master)
	}
}
