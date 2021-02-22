func isEvenOddTree(root *TreeNode) bool {
	q := make([]*TreeNode, 1, 16)
	q[0] = root
	d := 0
	for 0 != len(q) {
		v := 0
		if d&1 == 1 {
			v = 1000001
		}
		for l := len(q); l > 0; l-- {
			n := q[0]
			q = q[1:]
			if d&1 == 0 {
				if n.Val&1 == 0 || n.Val <= v {
					return false
				}
				v = n.Val
			} else {
				if n.Val&1 == 1 || n.Val >= v {
					return false
				}
				v = n.Val
			}
			if nil != n.Left {
				q = append(q, n.Left)
			}
			if nil != n.Right {
				q = append(q, n.Right)
			}
		}
		d++
	}
	return true
}
