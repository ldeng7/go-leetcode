import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	bs := []byte(strconv.Itoa(n))
	l := len(bs)
	i := l - 1
	for ; i > 0; i-- {
		if bs[i] > bs[i-1] {
			break
		}
	}
	if i == 0 {
		return -1
	}
	for j := l - 1; j >= i; j-- {
		if bs[j] > bs[i-1] {
			bs[j], bs[i-1] = bs[i-1], bs[j]
			break
		}
	}
	bs1 := bs[i:]
	sort.Slice(bs1, func(i, j int) bool {
		return bs1[i] < bs1[j]
	})
	copy(bs[i:], bs1)
	o, _ := strconv.Atoi(string(bs))
	if o > math.MaxInt32 {
		return -1
	}
	return o
}
