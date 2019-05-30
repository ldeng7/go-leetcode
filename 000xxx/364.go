func depthSumInverse(nestedList []*NestedInteger) int {
	u, w := 0, 0
	for 0 != len(nestedList) {
		next := []*NestedInteger{}
		for _, e := range nestedList {
			if e.IsInteger() {
				u += e.GetInteger()
			} else {
				next = append(next, e.GetList()...)
			}
		}
		w += u
		nestedList = next
	}
	return w
}
