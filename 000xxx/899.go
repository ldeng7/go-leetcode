import (
	"bytes"
	"sort"
)

func orderlyQueue(S string, K int) string {
	bs := []byte(S)
	if K != 1 {
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}
	bsm, bs0 := bs, bs
	for i := 0; i < len(bs0); i++ {
		bs = append(bs, bs0[i])
		bs = bs[1:]
		if bytes.Compare(bsm, bs) == 1 {
			bsm = bs
		}
	}
	return string(bsm)
}
