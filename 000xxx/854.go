import (
	"bytes"
	"math"
)

func kSimilarity(A string, B string) int {
	bsa, bsb := []byte(A), []byte(B)
	cache := map[string]int{}
	var cal func(int) int
	cal = func(i int) int {
		if i == len(bsa) || bytes.Compare(bsa, bsb) == 0 {
			return 0
		}
		a := string(bsa)
		if out, ok := cache[a]; ok {
			return out
		}
		out := math.MaxInt64
		if bsa[i] != bsb[i] {
			for j := i + 1; j < len(bsa); j++ {
				if bsa[j] == bsb[i] && bsa[j] != bsb[j] {
					bsa[i], bsa[j] = bsa[j], bsa[i]
					t := cal(i+1) + 1
					if t < out {
						out = t
					}
					bsa[i], bsa[j] = bsa[j], bsa[i]
				}
			}
		} else {
			out = cal(i + 1)
		}
		cache[a] = out
		return out
	}
	return cal(0)
}
