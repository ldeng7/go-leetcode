var cache = [22][]*TreeNode{}

func init() {
	for i := 2; i <= 20; i += 2 {
		cache[i] = []*TreeNode{}
	}
	cache[1] = []*TreeNode{&TreeNode{}}
	cache[3] = []*TreeNode{&TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}}
}

func allPossibleFBT(N int) []*TreeNode {
	var cal func(n int) []*TreeNode
	cal = func(n int) []*TreeNode {
		if nil != cache[n] {
			return cache[n]
		}
		out := []*TreeNode{}
		for i := 1; i <= n-2; i += 2 {
			for _, nl := range allPossibleFBT(i) {
				for _, nr := range allPossibleFBT(n - i - 1) {
					out = append(out, &TreeNode{Left: nl, Right: nr})
				}
			}
		}
		cache[n] = out
		return out
	}
	return cal(N)
}
