func nextGreatestLetter(letters []byte, target byte) byte {
	var o byte = 255
	for _, b := range letters {
		if b > target && b < o {
			o = b
		}
	}
	if o == 255 {
		for _, b := range letters {
			if b < o {
				o = b
			}
		}
	}
	return o
}
