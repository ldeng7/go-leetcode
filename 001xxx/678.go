func interpret(command string) string {
	l := len(command)
	bs := make([]byte, 0, l)
	for i := 0; i < l; {
		if command[i] == 'G' {
			bs = append(bs, 'G')
			i++
		} else if command[i+1] == ')' {
			bs = append(bs, 'o')
			i += 2
		} else {
			bs = append(bs, 'a', 'l')
			i += 4
		}
	}
	return string(bs)
}
