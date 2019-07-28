func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	o := []*TreeNode{}
	m := map[int]bool{}
	for _, i := range toDelete {
		m[i] = true
	}

	var cal func(**TreeNode, bool)
	cal = func(p **TreeNode, d bool) {
		n := *p
		if nil == n {
			return
		}
		d1 := false
		if m[n.Val] {
			d1, *p = true, nil
		} else if d {
			o = append(o, n)
		}
		cal(&n.Left, d1)
		cal(&n.Right, d1)
	}
	cal(&root, true)
	return o
}
