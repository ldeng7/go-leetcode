func shoppingOffers(price []int, special [][]int, needs []int) int {
	var cal func() int
	cal = func() int {
		o := 0
		for i, p := range price {
			o += p * needs[i]
		}
		for _, offer := range special {
			b := true
			for i, n := range needs {
				if n-offer[i] < 0 {
					b = false
				}
				needs[i] -= offer[i]
			}
			if b {
				o1 := cal() + offer[len(offer)-1]
				if o1 < o {
					o = o1
				}
			}
			for i, _ := range needs {
				needs[i] += offer[i]
			}
		}
		return o
	}
	return cal()
}
