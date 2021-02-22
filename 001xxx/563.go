import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func stoneGameV(stoneValue []int) int {
	l := len(stoneValue)
	ss := make([]int, l+1)
	for i := 1; i < l+1; i++ {
		ss[i] = ss[i-1] + stoneValue[i-1]
	}
	t := make([]int, l*l)
	for i := 0; i < l*l; i++ {
		t[i] = -1
	}

	var cal func(int, int) int
	cal = func(i, j int) int {
		k := i*l + j
		v := t[k]
		if v != -1 {
			return v
		} else if i == j {
			t[k] = 0
			return 0
		}
		s := (ss[i] + ss[j+1]) / 2
		h := sort.Search(j-i+1, func(k int) bool {
			return ss[i+k] >= s
		})
		for m := i + h; m <= j; m++ {
			a, b := ss[m]-ss[i], ss[j+1]-ss[m]
			if min(a, b)*2 < v {
				break
			}
			if a < b {
				v = max(v, cal(i, m-1)+a)
			} else if a > b {
				v = max(v, cal(m, j)+b)
			} else {
				v = max(v, max(cal(i, m-1), cal(m, j))+a)
			}
		}
		for m := i + h - 1; m >= i+1; m-- {
			a, b := ss[m]-ss[i], ss[j+1]-ss[m]
			if min(a, b)*2 < v {
				break
			}
			if a < b {
				v = max(v, cal(i, m-1)+a)
			} else if a > b {
				v = max(v, cal(m, j)+b)
			} else {
				v = max(v, max(cal(i, m-1), cal(m, j))+a)
			}
		}
		t[k] = v
		return v
	}

	return cal(0, l-1)
}
