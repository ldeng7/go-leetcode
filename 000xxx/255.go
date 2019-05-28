import "math"

func verifyPreorder(preorder []int) bool {
	l, i := math.MinInt64, -1
	for _, p := range preorder {
		if p < l {
			return false
		}
		for i >= 0 && p > preorder[i] {
			l, i = preorder[i], i-1
		}
		i++
		preorder[i] = p
	}
	return true
}
