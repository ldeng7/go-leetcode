import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	lh, lv := len(horizontalCuts), len(verticalCuts)
	hm, wm := h-horizontalCuts[lh-1], w-verticalCuts[lv-1]

	for i := 1; i < lh; i++ {
		hm = max(hm, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for i := 1; i < lv; i++ {
		wm = max(wm, verticalCuts[i]-verticalCuts[i-1])
	}
	hm, wm = max(hm, horizontalCuts[0]), max(wm, verticalCuts[0])
	return (hm * wm) % 1000000007
}
