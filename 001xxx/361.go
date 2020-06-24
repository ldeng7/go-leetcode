func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	c := 0
	for i := 0; i < n; i++ {
		if leftChild[i] == -1 {
			c++
		}
		if rightChild[i] == -1 {
			c++
		}
	}
	return c == n+1
}
