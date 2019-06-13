import (
	"sort"
	"strconv"
)

var m map[string]bool = map[string]bool{
	"1": true, "2": true, "4": true, "8": true, "16": true, "23": true, "46": true, "128": true, "256": true,
	"125": true, "0124": true, "0248": true, "0469": true, "1289": true, "13468": true, "23678": true, "35566": true,
	"011237": true, "122446": true, "224588": true, "0145678": true, "0122579": true, "0134449": true, "0368888": true,
	"11266777": true, "23334455": true, "01466788": true, "112234778": true, "234455668": true, "012356789": true,
}

func reorderedPowerOf2(N int) bool {
	bs := []byte(strconv.Itoa(N))
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return m[string(bs)]
}
