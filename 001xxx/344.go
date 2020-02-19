func angleClock(hour int, minutes int) float64 {
	a := float64(minutes)*5.5 - float64(hour)*30
	if a < 0 {
		a = -a
	}
	if a > 180 {
		return 360 - a
	}
	return a
}
