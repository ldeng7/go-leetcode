func slowestKey(releaseTimes []int, keysPressed string) byte {
	l := len(releaseTimes)
	for i := l - 1; i > 0; i-- {
		releaseTimes[i] -= releaseTimes[i-1]
	}
	t := 0
	for i := 1; i < l; i++ {
		if r, rt := releaseTimes[i], releaseTimes[t]; r > rt || (r == rt && keysPressed[i] > keysPressed[t]) {
			t = i
		}
	}
	return keysPressed[t]
}
