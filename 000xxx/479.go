import (
	"math"
	"strconv"
)

func largestPalindrome(n int) int {
	h := int(math.Pow(10, float64(n))) - 1
	l := h / 10
	for i := h; i > l; i-- {
		s := strconv.Itoa(i)
		bs := []byte(s)
		for j := len(s) - 1; j >= 0; j-- {
			bs = append(bs, s[j])
		}
		m, _ := strconv.Atoi(string(bs))
		for j := h; j*j > m; j-- {
			if m%j == 0 {
				return m % 1337
			}
		}
	}
	return 9
}
