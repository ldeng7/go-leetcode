func navigation(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		l, r := cal(n.Left), cal(n.Right)
		if nil != n.Left && nil != n.Right {
			if l == 0 && r == 0 {
				o++
				return 1
			} else if l == 0 || r == 0 {
				return 1
			}
			return 2
		} else if nil == n.Left {
			return r
		} else if nil == n.Right {
			return l
		}
		return 0
	}

	if cal(root.Left)+cal(root.Right) >= 2 {
		return o
	}
	return o + 1
}
