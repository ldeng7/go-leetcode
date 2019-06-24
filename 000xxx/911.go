import "sort"

type TopVotedCandidate struct {
	leaders [][2]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	votes, leaders := make([][2]int, len(persons)), [][2]int{}
	for i, p := range persons {
		votes[i] = [2]int{p, times[i]}
	}
	sort.Slice(votes, func(i, j int) bool { return votes[i][1] < votes[j][1] })
	m, c := map[int]int{}, 0
	for _, v := range votes {
		m[v[0]]++
		if m[v[0]] >= c {
			c = m[v[0]]
			leaders = append(leaders, v)
		}
	}
	return TopVotedCandidate{leaders}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.leaders)
	for l+1 < r {
		m := l + (r-l)>>1
		if t < this.leaders[m][1] {
			r = m
		} else {
			l = m
		}
	}
	return this.leaders[l][0]
}
