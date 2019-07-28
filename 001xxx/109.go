func corpFlightBookings(bookings [][]int, n int) []int {
	o := make([]int, n+1)
	for _, b := range bookings {
		o[b[0]-1] += b[2]
		o[b[1]] -= b[2]
	}
	for i := 1; i < n; i++ {
		o[i] += o[i-1]
	}
	return o[:n]
}
