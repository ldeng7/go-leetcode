import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func powerfulIntegers(x int, y int, bound int) []int {
	m := map[int]bool{}
	fx, fy, fb := float64(x), float64(y), float64(bound)
	ie, je := 20, 20
	if x > 1 {
		ie = min(20, int(math.Log(fb)/math.Log(fx)))
	}
	if y > 1 {
		je = min(20, int(math.Log(fb)/math.Log(fy)))
	}
	for i := 0; i <= ie; i++ {
		for j := 0; j <= je; j++ {
			k := int(math.Pow(fx, float64(i))) + int(math.Pow(fy, float64(j)))
			if k <= bound {
				m[k] = true
			}
		}
	}

	out, i := make([]int, len(m)), 0
	for k, _ := range m {
		out[i], i = k, i+1
	}
	return out
}
