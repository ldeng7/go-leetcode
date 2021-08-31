func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxValue(root *TreeNode, k int) int {
	arMax := func(ar []int) int {
		m := ar[0]
		for i := 1; i <= k; i++ {
			m = max(m, ar[i])
		}
		return m
	}

	var cal func(*TreeNode, []int)
	cal = func(n *TreeNode, ar []int) {
		if nil == n {
			return
		}
		arl, arr := [11]int{}, [11]int{}
		cal(n.Left, arl[:])
		cal(n.Right, arr[:])
		ar[0] = arMax(arl[:]) + arMax(arr[:])
		ar[1] = n.Val + arl[0] + arr[0]
		for i := 2; i <= k; i++ {
			m := 0
			for j := 0; j <= k && i-j-1 >= 0; j++ {
				m = max(m, arl[j]+arr[i-j-1]+n.Val)
			}
			ar[i] = m
		}
	}

	ar := [11]int{}
	cal(root, ar[:])
	return arMax(ar[:])
}
