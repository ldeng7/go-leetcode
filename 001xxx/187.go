import (
	"math"
	"sort"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	m, n := len(arr1), len(arr2)
	sort.Ints(arr2)
	set := make([]int, 1, n)
	set[0] = arr2[0]
	for i := 1; i < n; i++ {
		if arr2[i] != set[len(set)-1] {
			set = append(set, arr2[i])
		}
	}
	arr2, n = set, len(set)

	t := [2][]int{}
	t[0] = []int{arr1[0], arr2[0]}
	k, k1, lim := 1, 0, 0
	for i := 1; i < m; i++ {
		for j := lim; j < len(t[k1]); j++ {
			if j >= len(t[k]) {
				t[k] = append(t[k], math.MaxInt64)
			}
			if t[k1][j] < arr1[i] {
				t[k][j] = arr1[i]
			} else {
				t[k][j] = math.MaxInt64
			}
		}
		l := t[k1][lim]
		e := sort.Search(len(arr2), func(i int) bool { return arr2[i] > l })
		mi := min(e, n-1)
		for j := lim; j < len(t[k1]); j++ {
			for mi >= 0 && arr2[mi] > t[k1][j] {
				mi--
			}
			if mi+1 < n && arr2[mi+1] > t[k1][j] {
				if j+1 >= len(t[k]) {
					t[k] = append(t[k], math.MaxInt64)
				}
				t[k][j+1] = min(t[k][j+1], arr2[mi+1])
			}
			if t[k][j] == math.MaxInt64 {
				lim = j
			}
		}
		k, k1 = k^1, k1^1
	}
	if lim+1 >= len(t[k1]) || t[k1][lim+1] == math.MaxInt64 {
		return -1
	}
	return lim + 1
}
