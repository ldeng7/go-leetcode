func cal(n *TreeNode, d int) int {
	if nil == n {
		return 0
	} else if nil == n.Left && nil == n.Right && 0 == d {
		return 1
	} else if 0 == d {
		return 0
	}
	return cal(n.Left, d-1) + cal(n.Right, d-1)
}

func countPairs(root *TreeNode, distance int) int {
	if nil == root {
		return 0
	}
	l, r := root.Left, root.Right
	o := countPairs(l, distance) + countPairs(r, distance)
	if nil == l || nil == r {
		return o
	}
	distance -= 2
	for i := 0; i <= distance; i++ {
		for j := 0; j <= i; j++ {
			o += cal(l, j) * cal(r, i-j)
		}
	}
	return o
}
