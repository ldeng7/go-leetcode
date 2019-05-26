func closest(l, r int, m float64) int {
	if m-float64(l) < float64(r)-m {
		return l
	}
	return r
}

func closestValue(root *TreeNode, target float64) int {
	v := float64(root.Val)
	if target == v {
		return root.Val
	} else if target < v {
		if nil != root.Left {
			return closest(closestValue(root.Left, target), root.Val, target)
		}
		return root.Val
	} else {
		if nil != root.Right {
			return closest(root.Val, closestValue(root.Right, target), target)
		}
		return root.Val
	}
}
