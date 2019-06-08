import (
	"bytes"
	"strconv"
)

func optimalDivision(nums []int) string {
	buf := bytes.Buffer{}
	l := len(nums)
	for i, num := range nums {
		if i != 0 {
			buf.WriteByte('/')
		}
		if i == 1 && l > 2 {
			buf.WriteByte('(')
		}
		buf.WriteString(strconv.Itoa(num))
		if i == l-1 && l > 2 {
			buf.WriteByte(')')
		}
	}
	return buf.String()
}
