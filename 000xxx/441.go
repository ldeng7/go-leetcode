import "math"

func arrangeCoins(n int) int {
	return int((math.Sqrt(8.0*float64(n)+1.0) - 1.0) / 2)
}
