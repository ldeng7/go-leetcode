import "strconv"

func isAdditiveNumber(num string) bool {
	o, ar := false, []int{}
	var cal func(int)
	cal = func(b int) {
		if o {
			return
		}
		if b >= len(num) {
			if len(ar) >= 3 {
				o = true
			}
			return
		}
		for i := b; i < len(num); i++ {
			s := num[b : i+1]
			if (len(s) > 1 && s[0] == '0') || len(s) > 19 {
				break
			}
			n, _ := strconv.Atoi(s)
			if len(ar) >= 2 && n != ar[len(ar)-1]+ar[len(ar)-2] {
				continue
			}
			ar = append(ar, n)
			cal(i + 1)
			ar = ar[:len(ar)-1]
		}
	}
	cal(0)
	return o
}
