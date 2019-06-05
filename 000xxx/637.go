func averageOfLevels(root *TreeNode) []float64 {
	vals := [][]int{}
	var cal func(*TreeNode, int)
	cal = func(node *TreeNode, level int) {
		if nil == node {
			return
		}
		if level == len(vals) {
			vals = append(vals, []int{})
		}
		vals[level] = append(vals[level], node.Val)
		cal(node.Left, level+1)
		cal(node.Right, level+1)
	}
	cal(root, 0)
	out := make([]float64, len(vals))
	for i, ar := range vals {
		s := 0
		for _, val := range ar {
			s += val
		}
		out[i] = float64(s) / float64(len(ar))
	}
	return out
}
