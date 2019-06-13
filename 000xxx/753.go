import (
	"math"
	"strconv"
	"strings"
)

func crackSafe(n int, k int) string {
	out := strings.Repeat("0", n)
	m := map[string]bool{out: true}
	ie := int(math.Pow(float64(k), float64(n)))
	for i := 0; i < ie; i++ {
		sp := out[len(out)-n+1:]
		for j := k - 1; j >= 0; j-- {
			sj := strconv.Itoa(j)
			s := sp + sj
			if !m[s] {
				m[s] = true
				out += sj
				break
			}
		}
	}
	return out
}
