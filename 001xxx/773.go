func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	o, k := 0, 0
	switch ruleKey[0] {
	case 't':
		k = 0
	case 'c':
		k = 1
	case 'n':
		k = 2
	}
	for i := 0; i < len(items); i++ {
		if items[i][k] == ruleValue {
			o++
		}
	}
	return o
}
