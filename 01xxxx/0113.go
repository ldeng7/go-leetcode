func do(root *TreeNode, sum int, reg []int, out [][]int) [][]int {
	if nil == root {
		return out
	}
	reg = append(reg, root.Val)
	if nil == root.Left && nil == root.Right {
		if root.Val == sum {
			regNew := make([]int, len(reg))
			copy(regNew, reg)
			out = append(out, regNew)
		}
		return out
	}
	out = do(root.Left, sum-root.Val, reg, out)
	return do(root.Right, sum-root.Val, reg, out)
}

func pathSum(root *TreeNode, sum int) [][]int {
	out := [][]int{}
	return do(root, sum, []int{}, out)
}
