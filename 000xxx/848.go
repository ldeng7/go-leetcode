func shiftingLetters(S string, shifts []int) string {
	for i := len(shifts) - 2; i >= 0; i-- {
		shifts[i] = (shifts[i] + shifts[i+1]) % 26
	}
	bs := []byte(S)
	for i, shift := range shifts {
		bs[i] = byte((int(bs[i]-'a')+shift)%26) + 'a'
	}
	return string(bs)
}
