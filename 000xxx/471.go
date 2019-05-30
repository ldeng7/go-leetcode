import (
	"fmt"
	"strings"
)

func encode(s string) string {
	n := len(s)
	t := make([][]string, n)
	for i := 0; i < n; i++ {
		t[i] = make([]string, n)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j+i-1 < n; j++ {
			k := j + i - 1
			t[j][k] = s[j : j+i]
			sub, rep := s[j:k+1], ""
			pos := strings.Index(sub[1:]+sub, sub) + 1
			if pos < len(sub) {
				rep = fmt.Sprintf("%d[%s]", len(sub)/pos, t[j][j+pos-1])
				if len(rep) < len(t[j][k]) {
					t[j][k] = rep
				}
				continue
			}
			for l := j; l < k; l++ {
				left, right := t[j][l], t[l+1][k]
				if len(left)+len(right) < len(t[j][k]) {
					t[j][k] = left + right
				}
			}
		}
	}
	return t[0][n-1]
}
