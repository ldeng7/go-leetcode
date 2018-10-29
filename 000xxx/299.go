import "fmt"

func getHint(secret string, guess string) string {
	m := make([]int, 256)
	nb, nc := 0, 0
	for i, s := range []byte(secret) {
		if s == guess[i] {
			nb++
		} else {
			if m[s] < 0 {
				nc++
			}
			if m[guess[i]] > 0 {
				nc++
			}
			m[s]++
			m[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", nb, nc)
}
