func getCollisionTimes(cars [][]int) []float64 {
	l := len(cars)
	s := make([]int, 0, 16)
	o := make([]float64, l)
	for i := 0; i < l; i++ {
		o[i] = -1
	}
	for i := l - 1; i >= 0; i-- {
		pi, si := cars[i][0], cars[i][1]
		for 0 != len(s) && si <= cars[s[len(s)-1]][1] {
			s = s[:len(s)-1]
		}
		for 0 != len(s) {
			t := s[len(s)-1]
			o[i] = float64(cars[t][0]-pi) / float64(si-cars[t][1])
			if o[t] != -1 && o[i] > o[t] {
				s = s[:len(s)-1]
			} else {
				break
			}
		}
		s = append(s, i)
	}
	return o
}
