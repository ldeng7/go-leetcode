func cal(b, e int) []*TreeNode {
	out := []*TreeNode{}
	if b > e {
		return []*TreeNode{nil}
	}
	for i := b; i <= e; i++ {
		ls, rs := cal(b, i-1), cal(i+1, e)
		for _, l := range ls {
			for _, r := range rs {
				out = append(out, &TreeNode{i, l, r})
			}
		}
	}
	return out
}

func generateTrees(n int) []*TreeNode {
	if 0 == n {
		return []*TreeNode{}
	}
	return cal(1, n)
}
