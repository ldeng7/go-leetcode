func isValidSequence(root *TreeNode, arr []int) bool {
	l := len(arr)
	var cal func(*TreeNode, int) bool
	cal = func(n *TreeNode, i int) bool {
		if nil == n || n.Val != arr[i] {
			return false
		} else if i == l-1 {
			return n.Left == nil && n.Right == nil
		}
		return cal(n.Left, i+1) || cal(n.Right, i+1)
	}
	return cal(root, 0)
}
