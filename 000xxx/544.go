import (
	"fmt"
	"strconv"
)

func findContestMatch(n int) string {
	ss := make([]string, n)
	for i := 1; i <= n; i++ {
		ss[i-1] = strconv.Itoa(i)
	}
	for n > 1 {
		for i := 0; i < n>>1; i++ {
			ss[i] = fmt.Sprintf("(%s,%s)", ss[i], ss[n-i-1])
		}
		n >>= 1
	}
	return ss[0]
}
