func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for {
		if memory1 >= memory2 {
			if memory1 < i {
				break
			}
			memory1 -= i
		} else {
			if memory2 < i {
				break
			}
			memory2 -= i
		}
		i++
	}
	return []int{i, memory1, memory2}
}
