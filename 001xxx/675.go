import (
	"math/bits"
	"sort"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minimumDeviation(nums []int) int {
	ma := 1
	for _, n := range nums {
		ma = max(ma, n>>bits.TrailingZeros32(uint32(n)))
	}
	mi, mb := ma, bits.LeadingZeros32(uint32(ma))
	ar := make([]int, 0, len(nums))
	for _, n := range nums {
		if n&1 == 1 {
			n <<= 1
		}
		if n >= ma {
			n >>= mb - bits.LeadingZeros32(uint32(n))
			if n < ma {
				n <<= 1
			}
			ar = append(ar, n)
		}
		mi = min(mi, n)
	}

	sort.Ints(ar)
	o := ar[len(ar)-1] - mi
	for i := len(ar) - 1; ar[i] > ma; i-- {
		mi = min(mi, ar[i]>>1)
		o = min(o, ar[i-1]-mi)
	}
	return o
}
