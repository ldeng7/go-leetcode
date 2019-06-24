import (
	"math"
	"strconv"
)

func smallestGoodBase(str string) string {
	n, _ := strconv.Atoi(str)
	for i := int(math.Log2(float64(n + 1))); i >= 2; i-- {
		l, r := 2, int(math.Pow(float64(n), 1/float64(i-1)))+1
		for l < r {
			m, s := l+(r-l)>>1, 0
			for j := 0; j < i; j++ {
				s = s*m + 1
			}
			if s == n {
				return strconv.Itoa(m)
			} else if s < n {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return strconv.Itoa(n - 1)
}
