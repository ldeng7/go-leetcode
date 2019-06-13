import (
	"fmt"
	"math"
	"strconv"
)

func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func minimizeError(prices []string, target int) string {
	l := len(prices)
	lo, hi := make([]float64, l), make([]float64, l)
	t := make([][]float64, l+1)
	for i := 0; i <= l; i++ {
		t[i] = make([]float64, l+1)
	}

	sl, sh := 0, 0
	var slf, shf float64
	for i, sp := range prices {
		f, _ := strconv.ParseFloat(sp, 64)
		ff := math.Floor(f)
		lo[i], sl = f-ff, sl+int(ff)
		if f != ff {
			hi[i], sh = 1-lo[i], sh+int(ff)+1
		} else {
			hi[i], sh = 0, sh+int(ff)
		}
		slf, shf = slf+lo[i], shf+hi[i]
		t[i+1][0], t[i+1][i+1] = slf, shf
	}
	if target < sl || target > sh {
		return "-1"
	}

	for i := 1; i <= l; i++ {
		for j := 1; j < i; j++ {
			t[i][j] = min(t[i-1][j]+lo[i-1], t[i-1][j-1]+hi[i-1])
		}
	}
	return fmt.Sprintf("%.3f", t[l][target-sl])
}
