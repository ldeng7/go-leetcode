import "math"

func cal(node *TreeNode, pre, out *int) {
	if nil == node {
		return
	}
	cal(node.Left, pre, out)
	if -1 != *pre && node.Val-*pre < *out {
		*out = node.Val - *pre
	}
	*pre = node.Val
	cal(node.Right, pre, out)
}

func minDiffInBST(root *TreeNode) int {
	pre, out := -1, math.MaxInt64
	cal(root, &pre, &out)
	return out
}

