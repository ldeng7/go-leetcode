import "math"

func closestDivisors(num int) []int {
	n1, n2 := num+1, num+2
	sq := int(math.Sqrt(float64(n2)))
	for ; sq > 0; sq-- {
		if n1%sq == 0 {
			return []int{sq, n1 / sq}
		} else if n2%sq == 0 {
			return []int{sq, n2 / sq}
		}
	}
	return []int{0, 0}
}
