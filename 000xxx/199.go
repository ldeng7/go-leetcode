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
			node := q[len(q)-1]
			q = q[:len(q)-1]
			if nil != node.Left {
				q = append(q, nil)
				copy(q[1:], q[:len(q)-1])
				q[0] = node.Left
			}
			if nil != node.Right {
				q = append(q, nil)
				copy(q[1:], q[:len(q)-1])
				q[0] = node.Right
			}
		}
	}
	return out
}
