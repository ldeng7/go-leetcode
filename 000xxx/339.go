func depthSum(nestedList []*NestedInteger) int {
	var cal func([]*NestedInteger, int) int
	cal = func(ar []*NestedInteger, d int) int {
		sum := 0
		for _, i := range ar {
			if i.IsInteger() {
				sum += i.GetInteger() * d
			} else {
				sum += cal(i.GetList(), d+1)
			}
		}
		return sum
	}
	return cal(nestedList, 1)
}
