func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, s := range stones {
		sum += s
	}
	k := sum >> 1
	t := make([]int, k+1)
	for i := 0; i < len(stones); i++ {
		for j := k; j >= stones[i]; j-- {
			t[j] = max(t[j], t[j-stones[i]]+stones[i])
		}
	}
	return abs(sum - (t[k] << 1))
}
