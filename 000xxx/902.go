import (
	"math"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	ns := []byte(strconv.Itoa(N))
	ld, ln := len(D), len(ns)
	if ld == 0 || ln == 0 {
		return 0
	}
	ds := make([]byte, ld)
	for i, d := range D {
		ds[i] = d[0]
	}
	o := 0
	for i := ln; i > 1; i-- {
		o += int(math.Pow(float64(ld), float64(i-1)))
	}
	if ld < ln && ds[ld-1] < ns[0] {
		return o + int(math.Pow(float64(ld), float64(ln)))
	}

	for i, j := ld-1, 0; i >= 0 && j <= ln-1; {
		if ds[i] < ns[j] {
			o += (i + 1) * int(math.Pow(float64(ld), float64(ln-j-1)))
			break
		} else if ds[i] == ns[j] {
			if j == ln-1 {
				o += i + 1
				break
			}
			o += i * int(math.Pow(float64(ld), float64(ln-j-1)))
			i, j = ld-1, j+1
		} else {
			i--
		}
	}
	return o
}
