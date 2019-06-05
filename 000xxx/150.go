import "strconv"

func evalRPN(tokens []string) int {
	var cal func(*int) int
	cal = func(i *int) int {
		op := tokens[*i]
		if op != "+" && op != "-" && op != "*" && op != "/" {
			n, _ := strconv.Atoi(op)
			return n
		}
		*i--
		n1 := cal(i)
		*i--
		n2 := cal(i)
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
	i := len(tokens) - 1
	return cal(&i)
}
