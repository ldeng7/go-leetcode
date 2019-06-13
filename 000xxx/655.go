import "strconv"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func getHeight(n *TreeNode) int {
	if nil == n {
		return 0
	}
	return max(getHeight(n.Left), getHeight(n.Right)) + 1
}

func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	w := 1<<uint(h) - 1
	out := make([][]string, h)
	for i := 0; i < h; i++ {
		out[i] = make([]string, w)
	}

	var cal func(*TreeNode, int, int, int)
	cal = func(n *TreeNode, i, j, h1 int) {
		if nil == n || h1 == h {
			return
		}
		out[h1][(i+j)>>1] = strconv.Itoa(n.Val)
		cal(n.Left, i, (i+j)>>1, h1+1)
		cal(n.Right, (i+j)>>1+1, j, h1+1)
	}
	cal(root, 0, w-1, 0)
	return out
}
