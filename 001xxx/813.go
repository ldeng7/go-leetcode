func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	sentence1, sentence2 = " "+sentence1+" ", " "+sentence2+" "
	l1, l2 := len(sentence1), len(sentence2)
	if l1 < l2 {
		l1, l2 = l2, l1
		sentence1, sentence2 = sentence2, sentence1
	}
	l, r := 0, l2-1
	for i := 1; i < l2; i++ {
		if sentence1[i] != sentence2[i] {
			break
		}
		if sentence2[i] == ' ' {
			l = i
		}
	}
	for i := l2 - 1; i >= 0; i-- {
		if sentence1[l1-l2+i] != sentence2[i] {
			break
		}
		if sentence2[i] == ' ' {
			r = i
		}
	}
	return (l >= r)
}
