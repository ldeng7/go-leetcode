func game(guess []int, answer []int) int {
	o := 0
	for i := 0; i < 3; i++ {
		if guess[i] == answer[i] {
			o++
		}
	}
	return o
}
