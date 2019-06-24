func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	i, j := 0, 0
	for {
		if 0 != len(st) && st[len(st)-1] == popped[j] {
			st, j = st[:len(st)-1], j+1
		} else {
			if i >= len(pushed) {
				break
			}
			st, i = append(st, pushed[i]), i+1
		}
	}
	return 0 == len(st)
}
