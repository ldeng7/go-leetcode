import "math"

func consecutiveNumbersSum(N int) int {
	N <<= 1
	o := 0
	for i := int(math.Sqrt(float64(N))); i >= 1; i-- {
		if N%i == 0 && ((i+N/i)&1 == 1) {
			o++
		}
	}
	return o
}
