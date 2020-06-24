import "sort"

func rankTeams(votes []string) string {
	l := len(votes[0])
	cnts := make([][]int, 26)
	for i := 0; i < 26; i++ {
		ar := make([]int, l+1)
		ar[l] = -i
		cnts[i] = ar
	}
	for _, s := range votes {
		for i := 0; i < l; i++ {
			cnts[s[i]-'A'][i]++
		}
	}
	sort.Slice(cnts, func(i, j int) bool {
		ai, aj := cnts[i], cnts[j]
		for k := 0; k <= l; k++ {
			if ai[k] > aj[k] {
				return true
			} else if ai[k] < aj[k] {
				return false
			}
		}
		return true
	})

	n := l
	if n > 26 {
		n = 26
	}
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = 'A' + byte(-cnts[i][l])
	}
	return string(bs)
}
