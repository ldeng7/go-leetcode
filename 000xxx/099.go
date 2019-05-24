func recoverTree(root *TreeNode) {
	var f, s, p, n, prev *TreeNode = nil, nil, nil, root, nil
	for nil != n {
		if nil == n.Left {
			if nil != p && p.Val > n.Val {
				if nil == f {
					f = p
				}
				s = n
			}
			p, n = n, n.Right
		} else {
			prev = n.Left
			for nil != prev.Right && prev.Right != n {
				prev = prev.Right
			}
			if nil == prev.Right {
				prev.Right, n = n, n.Left
			} else {
				prev.Right = nil
				if p.Val > n.Val {
					if nil == f {
						f = p
					}
					s = n
				}
				p, n = n, n.Right
			}
		}
	}
	f.Val, s.Val = s.Val, f.Val
}
