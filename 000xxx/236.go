func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root || p == root || q == root {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	if nil != l && l != p && l != q {
		return l
	}
	r := lowestCommonAncestor(root.Right, p, q)
	if nil != l {
		if nil != r {
			return root
		}
		return l
	}
	return r
}
