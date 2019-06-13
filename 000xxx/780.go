func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx > ty {
			if ty == sy {
				return (tx-sx)%ty == 0
			}
			tx %= ty
		} else {
			if tx == sx {
				return (ty-sy)%tx == 0
			} else {
				ty %= tx
			}
		}
	}
	return tx == sx && ty == sy
}
