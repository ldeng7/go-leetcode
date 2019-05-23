func zigzagLevelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	if nil == root {
		return out
	}
	revNext := false
	nodes := []*TreeNode{root}
	for 0 != len(nodes) {
		vals := make([]int, len(nodes))
		if !revNext {
			for i, node := range nodes {
				vals[i] = node.Val
			}
		} else {
			for i, node := range nodes {
				vals[len(nodes)-i-1] = node.Val
			}
		}
		nodesNext := []*TreeNode{}
		for _, node := range nodes {
			if nil != node.Left {
				nodesNext = append(nodesNext, node.Left)
			}
			if nil != node.Right {
				nodesNext = append(nodesNext, node.Right)
			}
		}
		revNext = !revNext
		out = append(out, vals)
		nodes = nodesNext
	}
	return out
}
