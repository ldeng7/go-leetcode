import "math"

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	}
	n -= 5
	return int(math.Pow(3.0, float64(n/3+1))) * (n%3 + 2)
}
