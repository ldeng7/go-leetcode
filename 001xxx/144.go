func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func movesToMakeZigzag(nums []int) int {
	od, ev, l := 0, 0, len(nums)
	for i := 0; i < l; i += 2 {
		t := nums[i]
		if i > 0 {
			t = min(t, nums[i-1]-1)
		}
		if i+1 < l {
			t = min(t, nums[i+1]-1)
		}
		ev += nums[i] - t
	}
	for i := 1; i < l; i += 2 {
		t := nums[i]
		if i > 0 {
			t = min(t, nums[i-1]-1)
		}
		if i+1 < l {
			t = min(t, nums[i+1]-1)
		}
		od += nums[i] - t
	}
	return min(ev, od)
}
