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

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func findMinDifference(timePoints []string) int {
	m, ts := [1440]bool{}, []int{}
	for _, s := range timePoints {
		t := (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'0')*10 + int(s[4]-'0')
		if m[t] {
			return 0
		}
		m[t], ts = true, append(ts, t)
	}
	sort.Ints(ts)
	o, f, l, pt := math.MaxInt64, ts[0], ts[0], ts[0]
	for i := 1; i < len(ts); i++ {
		t := ts[i]
		o = min(o, t-pt)
		f, l, pt = min(f, t), max(l, t), t
	}
	return min(o, 1440+f-l)
}
