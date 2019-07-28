func pathSum(root *TreeNode, sum int) [][]int {
	o, t := [][]int{}, []int{}
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, sum int) {
		if nil == n {
			return
		}
		t = append(t, n.Val)
		if nil == n.Left && nil == n.Right && n.Val == sum {
			t1 := make([]int, len(t))
			copy(t1, t)
			o = append(o, t1)
		}
		cal(n.Left, sum-n.Val)
		cal(n.Right, sum-n.Val)
		t = t[:len(t)-1]
	}
	cal(root, sum)
	return o
}
