func change(amount int, coins []int) int {
	t := make([]int, amount+1)
	t[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			t[i] += t[i-c]
		}
	}
	return t[amount]
}
