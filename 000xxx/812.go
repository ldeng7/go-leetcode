func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func largestTriangleArea(points [][]int) float64 {
	var o float64
	for _, i := range points {
		for _, j := range points {
			for _, k := range points {
				a := float64(abs(i[0]*j[1]+j[0]*k[1]+k[0]*i[1]-j[0]*i[1]-k[0]*j[1]-i[0]*k[1])) * 0.5
				if a > o {
					o = a
				}
			}
		}
	}
	return o
}
