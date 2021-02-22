const MOD = 1000000007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(node *TreeNode, val int) *TreeNode {
	if nil == node {
		return &TreeNode{Val: val}
	} else if val < node.Val {
		node.Left = insert(node.Left, val)
	} else if val > node.Val {
		node.Right = insert(node.Right, val)
	}
	return node
}

func cal(node *TreeNode) int {
	if nil == node {
		return 0
	} else if nil == node.Left && nil == node.Right {
		return 1
	} else if nil == node.Left {
		return cal(node.Right)
	} else if nil == node.Right {
		return cal(node.Left)
	}

	l, r := cal(node.Left), cal(node.Right)
	var intl int
	n, k := node.Left.Val+node.Right.Val, node.Left.Val
	deg := func(n int) int {
		o := 0
		for u := MOD; u <= n; u *= MOD {
			o += n / u
		}
		return o
	}
	if deg(n)-deg(n-k) > deg(k) {
		intl = 0
	} else {
		cal := func(i, o int) int {
			for ; i%MOD == 0; i /= MOD {
			}
			return (o * i) % MOD
		}
		o := 1
		for i := n; i > n-k; i-- {
			o = cal(i, o)
		}
		d := 1
		for i := 1; i <= k; i++ {
			d = cal(i, d)
		}
		degree := 1
		for k := MOD - 2; k != 0; k >>= 1 {
			if k&1 == 1 {
				degree = (degree * d) % MOD
			}
			d = (d * d) % MOD
		}
		intl = (o * degree) % MOD
	}
	return (((l * r) % MOD) * (intl % MOD)) % MOD
}

func size(node *TreeNode) int {
	if nil == node {
		return 0
	}
	node.Val = size(node.Left) + size(node.Right) + 1
	return node.Val
}

func numOfWays(nums []int) int {
	var root *TreeNode
	for _, n := range nums {
		root = insert(root, n)
	}
	size(root)
	return (cal(root) - 1) % MOD
}
