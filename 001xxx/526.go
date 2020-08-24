func minNumberOperations(target []int) int {
	o, p := 0, 0
	for _, a := range target {
		if t := a - p; t > 0 {
			o += t
		}
		p = a
	}
	return o
}
