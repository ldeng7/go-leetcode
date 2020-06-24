func pseudoPalindromicPaths(root *TreeNode) int {
	o, ar := 0, [10]int{}
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, c int) {
		if nil == n {
			return
		}
		v := n.Val
		ar[v]++
		c += (ar[v]&1)*2 - 1
		if nil == n.Left && nil == n.Right {
			if c <= 1 {
				o++
			}
			ar[v]--
			return
		}
		cal(n.Left, c)
		cal(n.Right, c)
		ar[v]--
		return
	}
	cal(root, 0)
	return o
}
