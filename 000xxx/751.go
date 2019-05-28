import (
	"fmt"
	"strconv"
	"strings"
)

func ipToCIDR(ip string, n int) []string {
	out := []string{}
	es := strings.Split(ip, ".")
	var i uint64
	for j := 0; j < 4; j++ {
		i <<= 8
		e, _ := strconv.ParseUint(es[j], 10, 8)
		i |= e
	}

	un := uint64(n)
	for un != 0 {
		i1, d, t := i, uint64(1), 0
		for i1&1 == 0 {
			i1, d, t = i1>>1, d<<1, t+1
		}
		for d > un {
			d, t = d>>1, t-1
		}
		s := fmt.Sprintf("%d.%d.%d.%d/%d", i>>24, (i>>16)&0xff, (i>>8)&0xff, i&0xff, 32-t)
		out = append(out, s)
		i, un = i+d, un-d
	}
	return out
}
