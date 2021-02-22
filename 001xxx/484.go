func copyRandomBinaryTree(root *Node) *NodeCopy {
	m := map[*Node]*NodeCopy{}
	var cal func(*Node) *NodeCopy
	cal = func(n *Node) *NodeCopy {
		if nil == n {
			return nil
		} else if nc, ok := m[n]; ok {
			return nc
		}
		nc := &NodeCopy{Val: n.Val}
		m[n] = nc
		nc.Left = cal(n.Left)
		nc.Right = cal(n.Right)
		nc.Random = cal(n.Random)
		return nc
	}
	return cal(root)
}
