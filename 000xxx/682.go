import "strconv"

func calPoints(ops []string) int {
	i, ar := 0, make([]int, len(ops))
	for _, op := range ops {
		switch op {
		case "+":
			ar[i] = ar[i-1] + ar[i-2]
			i++
		case "D":
			ar[i] = ar[i-1] << 1
			i++
		case "C":
			ar[i-1] = 0
			i--
		default:
			ar[i], _ = strconv.Atoi(op)
			i++
		}
	}
	s := 0
	for _, i := range ar {
		s += i
	}
	return s
}
