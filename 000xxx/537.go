import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	la, lb := len(a), len(b)
	pa, pb := strings.LastIndexByte(a, '+'), strings.LastIndexByte(b, '+')
	ra, _ := strconv.Atoi(a[:pa])
	ia, _ := strconv.Atoi(a[pa+1 : la-1])
	rb, _ := strconv.Atoi(b[:pb])
	ib, _ := strconv.Atoi(b[pb+1 : lb-1])
	return fmt.Sprintf("%d+%di", ra*rb-ia*ib, ra*ib+ia*rb)
}
