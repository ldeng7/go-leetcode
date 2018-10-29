func asteroidCollision(asteroids []int) []int {
	out := []int{}
	for i := 0; i < len(asteroids); i++ {
		a := asteroids[i]
		if a > 0 {
			out = append(out, a)
		} else if len(out) == 0 || out[len(out)-1] < 0 {
			out = append(out, a)
		} else if -a > out[len(out)-1] {
			out = out[:len(out)-1]
			i--
		} else if -a == out[len(out)-1] {
			out = out[:len(out)-1]
		}
	}
	return out
}
