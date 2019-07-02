func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	o, i := []int{}, 0
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		if n.Val != voyage[i] {
			o, i = []int{-1}, i+1
			return
		}
		i++
		if i < len(voyage) && nil != n.Left && n.Left.Val != voyage[i] {
			o = append(o, n.Val)
			cal(n.Right)
			cal(n.Left)
		} else {
			cal(n.Left)
			cal(n.Right)
		}
	}

	cal(root)
	if 0 != len(o) && o[0] == -1 {
		return []int{-1}
	}
	return o
}
