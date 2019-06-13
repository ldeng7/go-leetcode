import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, line := range cpdomains {
		i := strings.IndexByte(line, ' ')
		c, _ := strconv.Atoi(line[:i])
		d := line[i+1:]

		m[d] += c
		i = strings.IndexByte(d, '.')
		m[d[i+1:]] += c
		j := strings.LastIndexByte(d, '.')
		if j != i {
			m[d[j+1:]] += c
		}
	}
	out, i := make([]string, len(m)), 0
	for d, c := range m {
		out[i], i = fmt.Sprintf("%d %s", c, d), i+1
	}
	return out
}
