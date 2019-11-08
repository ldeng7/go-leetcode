func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
	wm, wd, hm, hd := width%sideLength, width/sideLength, height%sideLength, height/sideLength
	d, m := wd*hd, wm*hm
	mm := wm
	if wd > hd {
		mm = hm
	}
	o1 := d*maxOnes + min(hm*sideLength, maxOnes)*wd + min(wm*sideLength, maxOnes)*hd + min(m, maxOnes)
	o2 := d*maxOnes + (m+maxOnes)*min(wd, hd) + min(mm*sideLength, maxOnes)*abs(wd-hd) + m
	return min(o1, o2)
}
