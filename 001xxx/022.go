func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, num int) {
		if nil == n {
			return
		}
		num = num<<1 + n.Val
		if nil == n.Left && nil == n.Right {
			sum += num
		} else {
			cal(n.Left, num)
			cal(n.Right, num)
		}
	}
	cal(root, 0)
	return sum
}
