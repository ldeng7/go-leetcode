func pathSum(root *TreeNode, sum int) [][]int {
	out := [][]int{}
	var cal func(*TreeNode, int, []int)
	cal = func(n *TreeNode, sum int, reg []int) {
		if nil == n {
			return
		}
		reg1 := make([]int, len(reg)+1)
		copy(reg1, reg)
		reg1[len(reg1)-1] = n.Val
		if nil == n.Left && nil == n.Right {
			if n.Val == sum {
				out = append(out, reg1)
			}
			return
		}
		cal(n.Left, sum-n.Val, reg1)
		cal(n.Right, sum-n.Val, reg1)
	}
	cal(root, sum, []int{})
	return out
}
