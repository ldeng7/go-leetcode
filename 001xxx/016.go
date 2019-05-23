import (
	"fmt"
	"strings"
)

func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		if -1 == strings.Index(S, fmt.Sprintf("%b", i)) {
			return false
		}
	}
	return true
}
