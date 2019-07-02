import "strconv"

func addOperators(num string, target int) []string {
	o := []string{}
	var cal func(string, int, int, string)
	cal = func(num string, diff, cur int, s string) {
		if 0 == len(num) && cur == target {
			o = append(o, s)
			return
		}
		for i := 1; i <= len(num); i++ {
			t := num[:i]
			if len(t) > 1 && t[0] == '0' {
				return
			}
			tn, _ := strconv.Atoi(t)
			num1 := num[i:]
			if 0 != len(s) {
				cal(num1, tn, cur+tn, s+"+"+t)
				cal(num1, -tn, cur-tn, s+"-"+t)
				cal(num1, diff*tn, (cur-diff)+diff*tn, s+"*"+t)
			} else {
				cal(num1, tn, tn, t)
			}
		}
	}
	cal(num, 0, 0, "")
	return o
}
