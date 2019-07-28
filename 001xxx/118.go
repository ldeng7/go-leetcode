var m = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func numberOfDays(Y int, M int) int {
	if M != 2 || Y&3 != 0 || (Y%100 == 0 && Y%400 != 0) {
		return m[M-1]
	}
	return 29
}
