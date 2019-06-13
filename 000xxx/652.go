import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	out, m := []*TreeNode{}, map[string]int{}
	var cal func(n *TreeNode) string
	cal = func(n *TreeNode) string {
		if nil == n {
			return "#"
		}
		s := fmt.Sprintf("%d,%s,%s", n.Val, cal(n.Left), cal(n.Right))
		if m[s] == 1 {
			out = append(out, n)
		}
		m[s]++
		return s
	}
	cal(root)
	return out
}
