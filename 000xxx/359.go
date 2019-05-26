type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{map[string]int{}}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if t, ok := this.m[message]; ok {
		if timestamp-t < 10 {
			return false
		}
	}
	this.m[message] = timestamp
	return true
}
