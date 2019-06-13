import "sort"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func numMovesStonesII(stones []int) []int {
	l := len(stones)
	sort.Ints(stones)
	mi, ma := 10001, max(stones[l-1]-stones[1], stones[l-2]-stones[0])-l+2
	j := 0
	for i, s := range stones {
		for ; j+1 < l && stones[j+1]-s < l; j++ {
		}
		k := l - (j - i + 1)
		if j-i+1 == l-1 && stones[j]-s+1 == l-1 {
			k = 2
		}
		if k < mi {
			mi = k
		}
	}
	return []int{mi, ma}
}
