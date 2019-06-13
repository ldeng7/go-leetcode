import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shortestToChar(S string, C byte) []int {
	indices := []int{}
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			indices = append(indices, i)
		}
	}
	l := len(indices)
	out := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] != C {
			k := sort.Search(l, func(j int) bool {
				return indices[j] >= i
			})
			if k == 0 {
				out[i] = indices[k] - i
			} else if k == l {
				out[i] = i - indices[l-1]
			} else {
				out[i] = min(i-indices[k-1], indices[k]-i)
			}
		} else {
			out[i] = 0
		}
	}
	return out
}
