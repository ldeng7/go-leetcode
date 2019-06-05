func coinChange(coins []int, amount int) int {
	t := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		t[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i {
				if t1 := t[i-c] + 1; t1 < t[i] {
					t[i] = t1
				}
			}
		}
	}
	if t[amount] > amount {
		return -1
	}
	return t[amount]
}
