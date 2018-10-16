func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if nil == root {
		return newNode
	}
	p := root
	var pn **TreeNode
	for {
		if val < p.Val {
			pn = &p.Left
		} else {
			pn = &p.Right
		}
		if nil == *pn {
			*pn = newNode
			break
		}
		p = *pn
	}
	return root
}
