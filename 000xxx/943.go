import (
	"bytes"
	"strings"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func shortestSuperstring(A []string) string {
	l, minLen := len(A), 241
	o, t, visited := make([]int, l), make([]int, 0, l), make([]bool, l)
	sufs := make([][]int, l)
	for i := 0; i < l; i++ {
		sufs[i] = make([]int, l)
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			a, b := A[i], A[j]
			k := min(len(a), len(b)) - 1
			for ; !strings.HasSuffix(a, b[:k]); k-- {
			}
			sufs[i][j] = k
		}
	}

	var cal func(int)
	cal = func(n int) {
		if len(t) == l {
			if minLen > n {
				minLen = n
				copy(o, t)
			}
			return
		}
		maxLen := -1
		ls := sufs[t[len(t)-1]]
		for i, sl := range ls {
			if sl > maxLen && !visited[i] {
				maxLen = sl
			}
		}
		for i, sl := range ls {
			if sl < maxLen || visited[i] {
				continue
			}
			t, visited[i] = append(t, i), true
			cal(n + len(A[i]) - maxLen)
			t, visited[i] = t[:len(t)-1], false
		}
	}

	for i := 0; i < l; i++ {
		t, visited[i] = append(t, i), true
		cal(len(A[i]))
		t, visited[i] = t[:len(t)-1], false
	}

	buf := &bytes.Buffer{}
	i := o[0]
	buf.WriteString(A[i])
	for k := 1; k < l; k++ {
		j := o[k]
		buf.WriteString(A[j][sufs[i][j]:])
		i = j
	}
	return buf.String()
}
