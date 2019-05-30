import "sort"

func verticalOrder(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}
	m := map[int][]int{}
	ari := []int{}
	qi := []int{0}
	qn := []*TreeNode{root}

	for 0 != len(qi) {
		it, nt := qi[0], qn[0]
		qi, qn = qi[1:], qn[1:]
		if _, ok := m[it]; !ok {
			ari = append(ari, it)
		}
		m[it] = append(m[it], nt.Val)
		if nil != nt.Left {
			qi, qn = append(qi, it-1), append(qn, nt.Left)
		}
		if nil != nt.Right {
			qi, qn = append(qi, it+1), append(qn, nt.Right)
		}
	}

	sort.Ints(ari)
	out := make([][]int, len(ari))
	for i, idx := range ari {
		out[i] = m[idx]
	}
	return out
}
