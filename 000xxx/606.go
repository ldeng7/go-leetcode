import (
	"strconv"
	"strings"
)

func tree2str(t *TreeNode) string {
	var sb strings.Builder
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		sb.WriteString(strconv.Itoa(n.Val))
		if nil == n.Left && nil == n.Right {
			return
		}
		sb.WriteByte('(')
		cal(n.Left)
		sb.WriteByte(')')
		if nil != n.Right {
			sb.WriteByte('(')
			cal(n.Right)
			sb.WriteByte(')')
		}
	}

	cal(t)
	return sb.String()
}
