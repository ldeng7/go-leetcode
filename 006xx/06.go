import "strconv"

func tree2str(t *TreeNode) string {
	if nil == t {
		return ""
	}
	s := strconv.Itoa(t.Val)
	if nil == t.Left && nil == t.Right {
		return s
	}
	s += "(" + tree2str(t.Left) + ")"
	if nil != t.Right {
		s += "(" + tree2str(t.Right) + ")"
	}
	return s
}
