import "sort"

type Node struct {
	node *TreeNode
	x    int
}

func verticalTraversal(root *TreeNode) [][]int {
	var ns []Node = []Node{{root, 0}}
	tp, tn := make([][][]int, 0, 16), make([][][]int, 0, 16)
	d := -1
	for 0 != len(ns) {
		d++
		ns1 := make([]Node, 0, 16)
		for _, n := range ns {
			node := n.node
			var i, j int
			var t *[][][]int
			if n.x >= 0 {
				i, j, t = n.x, (d-n.x)>>1, &tp
			} else {
				i, j, t = -n.x-1, (d+n.x)>>1, &tn
			}
			var ari *[][]int
			if i == len(*t) {
				*t = append(*t, make([][]int, 0, 16))
			}
			ari = &((*t)[i])
			var arj *[]int
			if j >= len(*ari) {
				for k := len(*ari); k <= j; k++ {
					*ari = append(*ari, make([]int, 0, 16))
				}
			}
			arj = &((*ari)[j])
			*arj = append(*arj, node.Val)

			if nil != node.Left {
				ns1 = append(ns1, Node{node.Left, n.x - 1})
			}
			if nil != node.Right {
				ns1 = append(ns1, Node{node.Right, n.x + 1})
			}
		}
		ns = ns1
	}

	o, oi := make([][]int, len(tn)+len(tp)), 0
	cal := func(ari [][]int) {
		c := 0
		for _, arj := range ari {
			c += len(arj)
		}
		ar, oj := make([]int, c), 0
		for _, arj := range ari {
			sort.Ints(arj)
			for _, v := range arj {
				ar[oj], oj = v, oj+1
			}
		}
		o[oi] = ar
	}
	for i := len(tn) - 1; i >= 0; i, oi = i-1, oi+1 {
		cal(tn[i])
	}
	for i := 0; i < len(tp); i, oi = i+1, oi+1 {
		cal(tp[i])
	}
	return o
}
