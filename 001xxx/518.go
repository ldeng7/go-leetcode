func numWaterBottles(numBottles int, numExchange int) int {
	return (numBottles*numExchange - 1) / (numExchange - 1)
}
