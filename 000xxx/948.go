import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	o, l, r := 0, 0, len(tokens)-1
	for l <= r {
		if tokens[l] > P {
			if o > 0 {
				P += tokens[r] - tokens[l]
				l, r = l+1, r-1
			} else {
				return o
			}
		} else {
			P -= tokens[l]
			l, o = l+1, o+1
		}
	}
	return o
}
