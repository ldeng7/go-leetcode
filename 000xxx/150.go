import "strconv"

func evalRPN(tokens []string) int {
	i := len(tokens) - 1
	var cal func() int
	cal = func() int {
		op := tokens[i]
		if op != "+" && op != "-" && op != "*" && op != "/" {
			n, _ := strconv.Atoi(op)
			return n
		}
		i--
		n1 := cal()
		i--
		n2 := cal()
		switch op {
		case "+":
			return n2 + n1
		case "-":
			return n2 - n1
		case "*":
			return n2 * n1
		case "/":
			return n2 / n1
		}
		return 0
	}
	return cal()
}
