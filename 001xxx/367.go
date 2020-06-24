func cal(ln *ListNode, tn *TreeNode) bool {
	if nil == ln {
		return true
	} else if nil == tn || ln.Val != tn.Val {
		return false
	}
	return cal(ln.Next, tn.Left) || cal(ln.Next, tn.Right)
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if nil == root {
		return false
	} else if cal(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}
