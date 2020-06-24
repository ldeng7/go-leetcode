func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minReorder(n int, connections [][]int) int {
	ar := make([][]int, n)
	for _, c := range connections {
		a, b := c[0], c[1]
		ar[a] = append(ar[a], b)
		ar[b] = append(ar[b], -a)
	}
	v := make([]bool, n)

	var cal func(int) int
	cal = func(i int) int {
		o := 0
		v[i] = true
		for _, j := range ar[i] {
			ja := abs(j)
			if !v[ja] {
				o += cal(ja)
				if j > 0 {
					o++
				}
			}
		}
		return o
	}
	return cal(0)
}
