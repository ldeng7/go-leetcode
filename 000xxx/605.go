func canPlaceFlowers(flowerbed []int, n int) bool {
	ar := make([]int, len(flowerbed)+2)
	copy(ar[1:], flowerbed)
	for i := 1; i < len(ar)-1; i++ {
		if ar[i] == 1 || ar[i-1] == 1 || ar[i+1] == 1 {
			continue
		}
		ar[i] = 1
		n--
	}
	return n <= 0
}
