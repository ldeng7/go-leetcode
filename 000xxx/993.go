func isCousins(root *TreeNode, x int, y int) bool {
	if x == root.Val || y == root.Val {
		return false
	}
	ns := []*TreeNode{root}
	for {
		ns1 := make([]*TreeNode, len(ns)<<1)
		for i, n := range ns {
			if nil != n {
				ns1[i<<1], ns1[i<<1+1] = n.Left, n.Right
			}
		}
		m := map[int]int{}
		for i, n := range ns1 {
			if nil != n {
				m[n.Val] = i
			}
		}
		if 0 == len(m) {
			break
		}
		ix, ex := m[x]
		iy, ey := m[y]
		if ex && ey {
			return ix&^1 != iy&^1
		} else if ex || ey {
			return false
		}
		ns = ns1
	}
	return false
}
