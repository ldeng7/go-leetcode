func busyStudent(startTime []int, endTime []int, queryTime int) int {
	o, l := 0, len(startTime)
	for i := 0; i < l; i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			o++
		}
	}
	return o
}
