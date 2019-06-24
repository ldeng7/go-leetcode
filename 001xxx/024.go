import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func videoStitching(clips [][]int, T int) int {
	t := make([]int, T+1)
	for i := 0; i <= T; i++ {
		t[i] = -1
	}
	sort.Slice(clips, func(i, j int) bool { return clips[i][0] < clips[j][0] })
	if clips[0][0] != 0 {
		return -1
	}
	for i := clips[0][0]; i <= min(clips[0][1], T); i++ {
		t[i] = 1
	}
	for i := 1; i < len(clips); i++ {
		b, e := clips[i][0], clips[i][1]
		if b > T {
			continue
		}
		if t[b] == -1 {
			return -1
		}
		for j := b; j <= e && j <= T; j++ {
			if b == 0 {
				t[j] = 1
			} else if t[j] == -1 {
				t[j] = t[b] + 1
			} else {
				t[j] = min(t[j], t[b]+1)
			}
		}
	}
	return t[T]
}
