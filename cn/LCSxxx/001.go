import "math/bits"

func leastMinutes(n int) int {
	return bits.Len(uint(n-1)) + 1
}
