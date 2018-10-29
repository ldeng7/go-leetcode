import "regexp"
import "strconv"
import "fmt"

func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func fractionAddition(expression string) string {
	re, _ := regexp.Compile("[+-]?\\d+")
	ss := re.FindAllString(expression, -1)
	nums := make([]int, len(ss))
	for i, s := range ss {
		n, _ := strconv.Atoi(s)
		nums[i] = n
	}
	a, b := 0, 1
	for i := 0; i < len(nums); i += 2 {
		a = a*nums[i+1] + b*nums[i]
		b *= nums[i+1]
		c := gcd(a, b)
		if c < 0 {
			c = -c
		}
		a /= c
		b /= c
	}
	return fmt.Sprintf("%d/%d", a, b)
}
