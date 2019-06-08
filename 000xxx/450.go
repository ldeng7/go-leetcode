func deleteNode(root *TreeNode, key int) *TreeNode {
	if nil == root {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if nil == root.Left {
			return root.Right
		}
		if nil == root.Right {
			return root.Left
		}
		n := root.Right
		for nil != n.Left {
			n = n.Left
		}
		root.Val = n.Val
		root.Right = deleteNode(root.Right, n.Val)
	}
	return root
}
