func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var cal func(*TreeNode, *[]int)
	cal = func(n *TreeNode, ls *[]int) {
		if nil == n {
			return
		}
		if nil == n.Left && nil == n.Right {
			*ls = append(*ls, n.Val)
			return
		}
		cal(n.Left, ls)
		cal(n.Right, ls)
	}
	ls1, ls2 := &[]int{}, &[]int{}
	cal(root1, ls1)
	cal(root2, ls2)
	if len(*ls1) != len(*ls2) {
		return false
	}
	for i := 0; i < len(*ls1); i++ {
		if (*ls1)[i] != (*ls2)[i] {
			return false
		}
	}
	return true
}
