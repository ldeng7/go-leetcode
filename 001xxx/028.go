func recoverFromPreorder(S string) *TreeNode {
	l, i := len(S), 0
	var cal func(int) *TreeNode
	cal = func(d0 int) *TreeNode {
		d, v := 0, 0
		for i < l && S[i] == '-' {
			d, i = d+1, i+1
		}
		if d != d0 {
			i -= d
			return nil
		}
		for i < l && S[i] != '-' {
			v, i = v*10+int(S[i]-'0'), i+1
		}
		return &TreeNode{v, cal(d0 + 1), cal(d0 + 1)}
	}
	return cal(0)
}
