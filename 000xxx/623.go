func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d <= 1 {
		n := &TreeNode{Val: v}
		if d == 1 {
			n.Left = root
		} else {
			n.Right = root
		}
		return n
	}
	if nil != root {
		d--
		if d > 1 {
			root.Left = addOneRow(root.Left, v, d)
			root.Right = addOneRow(root.Right, v, d)
		} else {
			root.Left = addOneRow(root.Left, v, 1)
			root.Right = addOneRow(root.Right, v, 0)
		}
	}
	return root
}
