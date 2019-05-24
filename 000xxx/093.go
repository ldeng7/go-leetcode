import "strconv"

func numStrLen(b int) int {
	if b < 10 {
		return 1
	} else if b < 100 {
		return 2
	}
	return 3
}

func restoreIpAddresses(s string) []string {
	out := []string{}
	var cal func(string, string, int)
	cal = func(si, so string, n int) {
		if 4 == n {
			if 0 == len(si) {
				out = append(out, so)
			}
			return
		}
		for i := 1; i < 4; i++ {
			if len(si) < i {
				break
			}
			if b, _ := strconv.Atoi(si[:i]); b >= 256 || numStrLen(b) != i {
				continue
			}
			so1 := so + si[:i]
			if n != 3 {
				so1 += "."
			}
			cal(si[i:], so1, n+1)
		}
	}
	cal(s, "", 0)
	return out
}
