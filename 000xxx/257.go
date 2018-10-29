import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	out := []string{}
	var cal func(*TreeNode, string)
	cal = func(node *TreeNode, s string) {
		if nil == node {
			return
		}
		s += fmt.Sprintf("->%d", node.Val)
		if nil == node.Left && nil == node.Right {
			out = append(out, s[2:])
			return
		}
		cal(node.Left, s)
		cal(node.Right, s)
	}
	cal(root, "")
	return out
}
