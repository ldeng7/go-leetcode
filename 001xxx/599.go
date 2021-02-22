func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	o, l := -2, len(customers)
	for i, cp, mp, cv := 0, 0, 0, 0; i < l || cv != 0; i++ {
		if i < l || cv < 4 {
			if i < l {
				cv += customers[i]
			}
			cb := min(cv, 4)
			cv -= cb
			cp += cb * boardingCost
			cp -= runningCost
		} else {
			cb := cv / 4 * 4
			cv -= cb
			cp += cb * boardingCost
			cp -= runningCost * (cb / 4)
			i += (cb / 4) - 1
		}
		if cp > mp {
			o, mp = i, cp
		}
	}
	return o + 1
}
