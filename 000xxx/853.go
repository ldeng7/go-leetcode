import "sort"

func carFleet(target int, position []int, speed []int) int {
	o, v0 := 0, 0.0
	m, s := map[int]float64{}, []int{}
	for i, p := range position {
		if _, ok := m[p]; !ok {
			s = append(s, p)
		}
		m[p] = float64(target-p) / float64(speed[i])
	}
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	for _, p := range s {
		if v := m[p]; v > v0 {
			v0, o = v, o+1
		}
	}
	return o
}
