import "sort"

func deckRevealedIncreasing(deck []int) []int {
	l := len(deck)
	if l <= 1 {
		return deck
	}
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))
	o := make([]int, 1, (l<<1)-2)
	o[0] = deck[0]
	for i := 1; i < l-1; i++ {
		o = append(o, deck[i], o[0])[1:]
	}
	o = append(o, deck[l-1])
	for i := 0; i < l>>1; i++ {
		o[i], o[l-i-1] = o[l-i-1], o[i]
	}
	return o
}
