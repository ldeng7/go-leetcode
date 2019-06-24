import "math"

func constructRectangle(area int) []int {
	r := int(math.Sqrt(float64(area)))
	for area%r != 0 {
		r--
	}
	return []int{area / r, r}
}
