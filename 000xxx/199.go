func rightSideView(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	q := []*TreeNode{root}
	out := []int{}
	for len(q) > 0 {
		out = append(out, q[0].Val)
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[len(q)-1]
			q = q[:len(q)-1]
			if nil != n.Left {
				q = append(q, nil)
				copy(q[1:], q)
				q[0] = n.Left
			}
			if nil != n.Right {
				q = append(q, nil)
				copy(q[1:], q)
				q[0] = n.Right
			}
		}
	}
	return out
}
