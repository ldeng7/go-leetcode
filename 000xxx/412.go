import "strconv"

func fizzBuzz(n int) []string {
	if n <= 0 {
		return []string{}
	}
	out := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			out[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			out[i-1] = "Fizz"
		} else if i%5 == 0 {
			out[i-1] = "Buzz"
		} else {
			out[i-1] = strconv.Itoa(i)
		}
	}
	return out
}
