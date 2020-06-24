import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSumOfLengths(arr []int, target int) int {
	o, p := math.MaxInt64, math.MaxInt64
	var l, h, c int
	q := make([]uint64, 0, 64)
	for h < len(arr) {
		c += arr[h]
		h++
		for c > target && l < h {
			c -= arr[l]
			l++
		}
		for 0 != len(q) {
			f := q[0]
			lf, hf := int(f&0xffffffff), int(f>>32)
			if hf > l {
				break
			}
			p = min(p, hf-lf)
			q = q[1:]
		}
		if c == target {
			if p != math.MaxInt64 {
				o = min(o, p+h-l)
			}
			q = append(q, uint64(l)|(uint64(h)<<32))
		}
	}
	if o == math.MaxInt64 {
		return -1
	}
	return o
}
