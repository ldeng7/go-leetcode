func lemonadeChange(bills []int) bool {
	f, t := 0, 0
	for _, b := range bills {
		if b == 5 {
			f++
		} else if b == 10 {
			f, t = f-1, t+1
		} else if t > 0 {
			f, t = f-1, t-1
		} else {
			f -= 3
		}
		if f < 0 {
			return false
		}
	}
	return true
}
