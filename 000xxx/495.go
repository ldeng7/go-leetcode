func findPoisonedDuration(timeSeries []int, duration int) int {
	o, l := 0, len(timeSeries)
	if 0 == l {
		return o
	}
	for i := 1; i < l; i++ {
		d := timeSeries[i] - timeSeries[i-1]
		if d < duration {
			o += d
		} else {
			o += duration
		}
	}
	return o + duration
}
