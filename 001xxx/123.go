func cal(n *TreeNode, d int) (int, *TreeNode) {
	if nil == n {
		return -1, nil
	}
	dl, l := cal(n.Left, d+1)
	dr, r := cal(n.Right, d+1)
	if nil == l && nil == r {
		return d + 1, n
	}
	if dl == dr {
		return dl, n
	} else if dl > dr {
		return dl, l
	}
	return dr, r
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, n := cal(root, 0)
	return n
}

