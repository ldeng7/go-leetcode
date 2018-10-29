var ps = map[int]bool{
	2: true, 3: true, 5: true, 7: true, 11: true, 13: true, 17: true, 19: true,
}

func countPrimeSetBits(L int, R int) int {
	out := 0
	for i := L; i <= R; i++ {
		j, c := i, 0
		for j > 0 {
			if j&1 == 1 {
				c++
			}
			j >>= 1
		}
		if ps[c] {
			out++
		}
	}
	return out
}
