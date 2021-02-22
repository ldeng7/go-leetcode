func cloneTree(root *Node) *Node {
	if nil == root {
		return nil
	}
	n := &Node{Val: root.Val, Children: make([]*Node, len(root.Children))}
	for i, c := range root.Children {
		n.Children[i] = cloneTree(c)
	}
	return n
}
