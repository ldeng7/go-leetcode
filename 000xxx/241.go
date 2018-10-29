import "strconv"

func diffWaysToCompute(s string) []int {
	out := []int{}
	for i, b := range []byte(s) {
		if b != '+' && b != '-' && b != '*' {
			continue
		}
		sll := diffWaysToCompute(s[:i])
		slr := diffWaysToCompute(s[i+1:])
		for _, il := range sll {
			for _, ir := range slr {
				switch b {
				case '+':
					out = append(out, il+ir)
				case '-':
					out = append(out, il-ir)
				case '*':
					out = append(out, il*ir)
				}
			}
		}
	}
	if 0 == len(out) {
		i, _ := strconv.Atoi(s)
		out = []int{i}
	}
	return out
}
