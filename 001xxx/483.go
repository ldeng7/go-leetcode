const L = 15

type TreeAncestor struct {
	t [][16]int
}

func Constructor(n int, parent []int) TreeAncestor {
	t := make([][L + 1]int, n)
	for i := 0; i < n; i++ {
		t[i][0] = parent[i]
		for j := 1; j <= L; j++ {
			t[i][j] = -1
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= L; j++ {
			if p := t[i][j-1]; p != -1 {
				t[i][j] = t[p][j-1]
			}
		}
	}
	return TreeAncestor{t}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	o, i := node, L
	for 0 != k {
		m := 1 << i
		for ; m > k; m, i = m>>1, i-1 {
		}
		o = this.t[o][i]
		k &^= m
		if o == -1 {
			return -1
		}
	}
	return o
}
