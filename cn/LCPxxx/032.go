import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Node struct {
	l, r, w int
}

func processTasks(tasks [][]int) int {
	ar := make([]Node, len(tasks))
	for i, t := range tasks {
		ar[i] = Node{t[0], t[1] + 1, t[2]}
	}
	sort.Slice(ar, func(i, j int) bool {
		return ar[i].r < ar[j].r
	})
	ar[0].l = ar[0].r - ar[0].w
	k := 1
	for i := 1; i < len(tasks); i++ {
		j := sort.Search(k, func(j int) bool {
			return ar[j].r > ar[i].l
		})
		e := ar[i].w
		if j < k {
			e -= ar[k-1].w - ar[j].w + ar[j].r - max(ar[j].l, ar[i].l)
		}
		if e <= 0 {
			continue
		}
		r := ar[i].r
		d := min(r-ar[k-1].r, e)
		l := r - d
		e -= d
		for e != 0 {
			k--
			l = ar[k].l
			if k > 0 {
				d = min(l-ar[k-1].r, e)
			} else {
				d = min(l, e)
			}
			l, e = l-d, e-d
		}
		if k > 0 {
			ar[k] = Node{l, r, ar[k-1].w + r - l}
		} else {
			ar[k] = Node{l, r, r - l}
		}
		k++
	}
	return ar[k-1].w
}
