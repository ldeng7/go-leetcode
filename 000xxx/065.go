const (
	A_INVALID = iota
	A_SPACE
	A_SIGN
	A_DIGIT
	A_POINT
	A_EXP
)

var trans [9][6]int = [...][6]int{
	{-1, 0, 3, 1, 2, -1},
	{-1, 8, -1, 1, 4, 5},
	{-1, -1, -1, 4, -1, -1},
	{-1, -1, -1, 1, 2, -1},
	{-1, 8, -1, 4, -1, 5},
	{-1, -1, 6, 7, -1, -1},
	{-1, -1, -1, 7, -1, -1},
	{-1, 8, -1, 7, -1, -1},
	{-1, 8, -1, -1, -1, -1},
}

var endStateValid [9]bool = [...]bool{false, true, false, false, true, false, false, true, true}

func isNumber(s string) bool {
	st := 0
	for _, c := range []byte(s) {
		act := A_INVALID
		switch c {
		case ' ':
			act = A_SPACE
		case '+', '-':
			act = A_SIGN
		case '.':
			act = A_POINT
		case 'e', 'E':
			act = A_EXP
		}
		if c >= '0' && c <= '9' {
			act = A_DIGIT
		}
		st = trans[st][act]
		if -1 == st {
			return false
		}
	}
	return endStateValid[st]
}
