import (
	"math"
	"strconv"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func nearestPalindromic(n string) string {
	o, l, m := 0, len(n), math.MaxInt64
	num, _ := strconv.Atoi(n)
	pref, _ := strconv.Atoi(n[:(l+1)>>1])
	s := map[int]bool{}
	s[int(math.Pow(10, float64(l)))+1] = true
	s[int(math.Pow(10, float64(l-1)))-1] = true
	for i := -1; i <= 1; i++ {
		bs := []byte(strconv.Itoa(pref + i))
		for j := len(bs) - 1 - (l & 1); j >= 0; j-- {
			bs = append(bs, bs[j])
		}
		k, _ := strconv.Atoi(string(bs))
		s[k] = true
	}
	delete(s, num)
	for i, _ := range s {
		d := abs(i - num)
		if d < m {
			m, o = d, i
		} else if (d == m) && i < o {
			o = i
		}
	}
	return strconv.Itoa(o)
}
