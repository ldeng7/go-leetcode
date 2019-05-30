func sequenceReconstruction(org []int, seqs [][]int) bool {
	if 0 == len(seqs) {
		return false
	}
	n, c, e := len(org), len(org)-1, false
	pos, flags := make([]int, n+1), make([]bool, n+1)
	for i := 0; i < n; i++ {
		pos[org[i]] = i
	}
	for _, seq := range seqs {
		for i, s := range seq {
			e = true
			if seq[i] <= 0 || seq[i] > n {
				return false
			}
			if i == 0 {
				continue
			}
			p := seq[i-1]
			if pos[p] >= pos[s] {
				return false
			}
			if (!flags[s]) && pos[p]+1 == pos[s] {
				flags[s], c = true, c-1
			}
		}
	}
	return c == 0 && e
}
