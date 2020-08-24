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

func n(m int) int {
	if m >= 100 {
		return 3
	} else if m >= 10 {
		return 2
	} else if m >= 2 {
		return 1
	}
	return 0
}

const N = 127

var t [N][N]int

func getLengthOfOptimalCompression(s string, k int) int {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			t[i][j] = -1
		}
	}
	l := len(s)
	var cal func(int, int) int
	cal = func(i, k int) int {
		if k < 0 {
			return N
		} else if i >= l-k {
			return 0
		}
		o := &(t[i][k])
		if -1 != *o {
			return *o
		}
		*o = N
		cnts := [26]int{}
		for j, m := i, 0; j < l; j++ {
			c := int(s[j] - 'a')
			cnts[c]++
			m = max(m, cnts[c])
			*o = min(*o, cal(j+1, k-(j-i-m+1))+n(m)+1)
		}
		return *o
	}
	return cal(0, k)
}
