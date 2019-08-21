func maxLevelSum(root *TreeNode) int {
	m, ml, l := 0, 0, 1
	var cal func([]*TreeNode)
	cal = func(ar []*TreeNode) {
		t := make([]*TreeNode, 0, len(ar)<<1)
		s := 0
		for _, n := range ar {
			if nil != n {
				t = append(t, n.Left, n.Right)
				s += n.Val
			}
		}
		if m < s {
			m, ml = s, l
		}
		if 0 != len(t) {
			l++
			cal(t)
		}
	}
	cal([]*TreeNode{root})
	return ml
}
