func findClosestLeaf(root *TreeNode, k int) int {
	m := map[*TreeNode]*TreeNode{}
	var cal func(n *TreeNode) *TreeNode
	cal = func(n *TreeNode) *TreeNode {
		if n.Val == k {
			return n
		}
		if nil != n.Left {
			m[n.Left] = n
			l := cal(n.Left)
			if nil != l {
				return l
			}
		}
		if nil != n.Right {
			m[n.Right] = n
			r := cal(n.Right)
			if nil != r {
				return r
			}
		}
		return nil
	}

	n := cal(root)
	q := []*TreeNode{n}
	s := map[*TreeNode]bool{n: true}
	for 0 != len(q) {
		n := q[0]
		q = q[1:]
		if nil == n.Left && nil == n.Right {
			return n.Val
		}
		if nil != n.Left && !s[n.Left] {
			s[n.Left], q = true, append(q, n.Left)
		}
		if nil != n.Right && !s[n.Right] {
			s[n.Right], q = true, append(q, n.Right)
		}
		if nil != m[n] && !s[m[n]] {
			s[m[n]], q = true, append(q, m[n])
		}
	}
	return -1
}
