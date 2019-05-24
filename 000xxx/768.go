func maxChunksToSorted(arr []int) int {
	st := []int{}
	for _, a := range arr {
		if 0 == len(st) || a >= st[len(st)-1] {
			st = append(st, a)
		} else {
			s := st[len(st)-1]
			st = st[:len(st)-1]
			for 0 != len(st) && a < st[len(st)-1] {
				st = st[:len(st)-1]
			}
			st = append(st, s)
		}
	}
	return len(st)
}
