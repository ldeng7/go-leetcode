func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minOperations(logs []string) int {
	o := 0
	for _, s := range logs {
		switch s {
		case "./":
			continue
		case "../":
			o = max(0, o-1)
		default:
			o++
		}
	}
	return o
}
