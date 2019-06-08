func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	out := []int{}
	var cal func(*TreeNode, int) int

	cal = func(n *TreeNode, d int) int {
		if nil == n {
			return 0
		}
		if d == K {
			out = append(out, n.Val)
			return 0
		}
		if n.Val == target.Val || d > 0 {
			d++
		}
		l, r := cal(n.Left, d), cal(n.Right, d)
		if l == K || r == K {
			out = append(out, n.Val)
			return 0
		}
		if n.Val == target.Val {
			return 1
		}
		if l > 0 {
			cal(n.Right, l+1)
		}
		if r > 0 {
			cal(n.Left, r+1)
		}
		if l > 0 {
			return l + 1
		} else if r > 0 {
			return r + 1
		}
		return 0
	}

	cal(root, 0)
	return out
}
