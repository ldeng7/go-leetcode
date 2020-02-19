import "sort"

func arrayRankTransform(arr []int) []int {
	l := len(arr)
	if 0 == l {
		return nil
	}
	arr1 := make([]int, l)
	copy(arr1, arr)
	sort.Ints(arr1)

	p, n := arr1[0], 1
	m := map[int]int{p: n}
	for i := 1; i < l; i++ {
		if a := arr1[i]; a != p {
			p, n = a, n+1
			m[a] = n
		}
	}

	for i, a := range arr {
		arr[i] = m[a]
	}
	return arr
}
