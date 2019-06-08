import (
	"math"
	"strconv"
)

func splitIntoFibonacci(S string) []int {
	var out, t []int
	var cal func(b int)
	cal = func(b int) {
		if nil != out {
			return
		}
		if b >= len(S) && len(t) >= 3 {
			out = make([]int, len(t))
			copy(out, t)
			return
		}
		for i := b; i < len(S); i++ {
			s := S[b : i+1]
			if (len(s) > 1 && s[0] == '0') || len(s) > 10 {
				break
			}
			n, _ := strconv.Atoi(s)
			if n > math.MaxInt32 {
				break
			}
			l := len(t)
			if l >= 2 && n != t[l-1]+t[l-2] {
				continue
			}
			t = append(t, n)
			cal(i + 1)
			t = t[:len(t)-1]
		}
	}
	cal(0)
	return out
}
