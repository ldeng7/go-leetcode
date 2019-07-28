import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := map[int]bool{}
	for _, i := range arr2 {
		m[i] = true
	}
	t := make([]int, 0, len(arr1))
	c := map[int]int{}
	for _, i := range arr1 {
		if m[i] {
			c[i]++
		} else {
			t = append(t, i)
		}
	}
	i := 0
	for _, j := range arr2 {
		for k := 0; k < c[j]; k++ {
			arr1[i], i = j, i+1
		}
	}
	sort.Ints(t)
	copy(arr1[i:], t)
	return arr1
}
