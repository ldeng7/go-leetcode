import (
	"sort"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	ws := make([]string, len(website))
	copy(ws, website)
	sort.Strings(ws)
	id, addr2Id, id2Addr := 1, map[string]int{}, map[int]string{}
	for _, w := range ws {
		if _, ok := addr2Id[w]; !ok {
			addr2Id[w], id2Addr[id] = id, w
			id++
		}
	}

	data := map[string][][2]int{}
	for i, u := range username {
		data[u] = append(data[u], [2]int{timestamp[i], addr2Id[website[i]]})
	}
	for name, ar := range data {
		sort.Slice(ar, func(i, j int) bool {
			return ar[i][0] < ar[j][0]
		})
		data[name] = ar
	}
	cnts := map[uint32]int{}
	for _, ar := range data {
		l := len(ar)
		set := map[uint32]bool{}
		for i := 0; i < l; i++ {
			idI := uint32(ar[i][1])
			for j := i + 1; j < l; j++ {
				idJ := uint32(ar[j][1])
				for k := j + 1; k < l; k++ {
					set[(idI<<16)|(idJ<<8)|uint32(ar[k][1])] = true
				}
			}
		}
		for k, _ := range set {
			cnts[k]++
		}
	}

	km, cm := uint32(0), 0
	for k, c := range cnts {
		if c > cm || (c == cm && k < km) {
			km, cm = k, c
		}
	}
	return []string{id2Addr[int(km>>16)], id2Addr[int(km>>8&0xff)], id2Addr[int(km&0xff)]}
}
