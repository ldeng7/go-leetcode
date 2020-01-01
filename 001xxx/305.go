func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	o := make([]int, 0, 32)
	var cal func(*[]*TreeNode, *TreeNode)
	cal = func(s *[]*TreeNode, n *TreeNode) {
		if nil == n {
			return
		}
		*s = append(*s, n)
		cal(s, n.Left)
	}

	sl1, sl2 := make([]*TreeNode, 0, 32), make([]*TreeNode, 0, 32)
	s1, s2 := &sl1, &sl2
	cal(s1, root1)
	cal(s2, root2)
	for 0 != len(*s1) || 0 != len(*s2) {
		var s *[]*TreeNode
		l1, l2 := len(*s1), len(*s2)
		if 0 != l1 && 0 != l2 {
			if (*s1)[l1-1].Val < (*s2)[l2-1].Val {
				s = s1
			} else {
				s = s2
			}
		} else if 0 == l2 {
			s = s1
		} else {
			s = s2
		}
		l := len(*s)
		n := (*s)[l-1]
		*s = (*s)[:l-1]
		o = append(o, n.Val)
		cal(s, n.Right)
	}
	return o
}
