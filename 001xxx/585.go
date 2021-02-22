func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func isTransformable(s string, t string) bool {
	bss, bst, l := []byte(s), []byte(t), len(s)
	ar := [10]int{}

	var cal func(int, int) bool
	cal = func(i, j int) bool {
		for j < l {
			if b := bss[i]; b == 0 {
				i++
			} else if b == bst[j] {
				i, j = i+1, j+1
			} else {
				break
			}
		}
		if j == l {
			return true
		}

		k, bj := max(i, ar[bst[j]-'0']), bst[j]
		for ; k < l; k++ {
			if b := bss[k]; b == 0 {
				continue
			} else if b == bj {
				bss[k] = 0
				ar[bj-'0'] = k + 1
				break
			} else if b < bj {
				return false
			}
		}
		if k == l {
			return false
		}
		return cal(i, j+1)
	}
	return cal(0, 0)
}
