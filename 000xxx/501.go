func findMode(root *TreeNode) []int {
	out := []int{}
	cnt, max := 1, 0
	var prev *TreeNode
	var iter func(*TreeNode)
	iter = func(node *TreeNode) {
		if nil == node {
			return
		}
		iter(node.Left)
		if nil != prev {
			if node.Val == prev.Val {
				cnt++
			} else {
				cnt = 1
			}
		}
		if cnt > max {
			out = []int{node.Val}
			max = cnt
		} else if cnt == max {
			out = append(out, node.Val)
			max = cnt
		}
		prev = node
		iter(node.Right)
	}
	iter(root)
	return out
}
