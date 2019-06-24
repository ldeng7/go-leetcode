func numberOfBoomerangs(points [][]int) int {
	o := 0
	for i := 0; i < len(points); i++ {
		m := map[int]int{}
		for j := 0; j < len(points); j++ {
			x, y := points[i][0]-points[j][0], points[i][1]-points[j][1]
			m[x*x+y*y]++
		}
		for _, c := range m {
			o += c * (c - 1)
		}
	}
	return o
}
