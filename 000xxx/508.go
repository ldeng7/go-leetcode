func findFrequentTreeSum(root *TreeNode) []int {
	o, m, c := []int{}, map[int]int{}, 0
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		s := n.Val + cal(n.Left) + cal(n.Right)
		m[s]++
		if m[s] >= c {
			if m[s] > c {
				o = []int{}
			}
			o, c = append(o, s), m[s]
		}
		return s
	}
	cal(root)
	return o
}
