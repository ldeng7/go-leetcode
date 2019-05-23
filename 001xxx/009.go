import "fmt"

func bitwiseComplement(N int) int {
	s := fmt.Sprintf("%b", N)
	i := 0
	for _, b := range []byte(s) {
		i <<= 1
		if '0' == b {
			i |= 1
		}
	}
	return i
}
