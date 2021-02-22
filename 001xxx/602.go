func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	q := make([]*TreeNode, 1, 16)
	q[0] = root
	for 0 != len(q) {
		ar := make([]*TreeNode, 0, 16)
		for i, n := range q {
			if n == u {
				if i == len(q)-1 {
					return nil
				}
				return q[i+1]
			}
			if nil != n.Left {
				ar = append(ar, n.Left)
			}
			if nil != n.Right {
				ar = append(ar, n.Right)
			}
		}
		q = ar
	}
	return nil
}
