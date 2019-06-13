import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	out, st, pt := make([]int, n), []int{}, 0
	for _, log := range logs {
		i1, i2 := strings.IndexByte(log, ':'), strings.LastIndexByte(log, ':')
		t, _ := strconv.Atoi(log[i2+1:])
		if 0 != len(st) {
			out[st[len(st)-1]] += t - pt
		}
		pt = t
		if log[i1+1:i2] == "start" {
			i, _ := strconv.Atoi(log[:i1])
			st = append(st, i)
		} else {
			i := st[len(st)-1]
			st = st[:len(st)-1]
			out[i], pt = out[i]+1, pt+1
		}
	}
	return out
}
