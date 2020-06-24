func maxScore(cardPoints []int, k int) int {
	l := len(cardPoints)
	a, b := 0, 0
	for i := 0; i < k; i++ {
		a += cardPoints[i]
	}
	o := a
	for i := 0; i < k; i++ {
		a -= cardPoints[k-i-1]
		b += cardPoints[l-i-1]
		if t := a + b; t > o {
			o = t
		}
	}
	return o
}
