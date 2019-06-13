import "strconv"

func solveEquation(equation string) string {
	l, a, b, s, j := len(equation), 0, 0, 1, 0
	for i := 0; i < l; i++ {
		switch equation[i] {
		case '+', '-':
			if i > j {
				n, _ := strconv.Atoi(equation[j:i])
				b += n * s
			}
			j = i
		case 'x':
			if i == j || equation[i-1] == '+' {
				a += s
			} else if equation[i-1] == '-' {
				a -= s
			} else {
				n, _ := strconv.Atoi(equation[j:i])
				a += n * s
			}
			j = i + 1
		case '=':
			if i > j {
				n, _ := strconv.Atoi(equation[j:i])
				b += n * s
			}
			s = -1
			j = i + 1
		}
	}
	if j < l {
		n, _ := strconv.Atoi(equation[j:])
		b += n * s
	}
	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(-b/a)
}
