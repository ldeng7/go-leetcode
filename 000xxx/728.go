func selfDividingNumbers(left int, right int) []int {
	out := []int{}
	for i := left; i <= right; i++ {
		if i <= 9 {
			out = append(out, i)
		} else {
			j := i
			for j > 0 {
				d := j % 10
				if 0 == d || i%(d) != 0 {
					break
				}
				j /= 10
			}
			if 0 == j {
				out = append(out, i)
			}
		}
	}
	return out
}
