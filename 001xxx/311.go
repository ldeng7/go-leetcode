import "sort"

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	q, l, set := []int{id}, 1, map[int]bool{id: true}
	for l != 0 && level != 0 {
		level--
		for l != 0 {
			l--
			ids := friends[q[0]]
			for _, id := range ids {
				if !set[id] {
					q = append(q, id)
					set[id] = true
				}
			}
			q = q[1:]
		}
		l = len(q)
	}

	cnts := map[string]int{}
	for _, id := range q {
		for _, s := range watchedVideos[id] {
			cnts[s]++
		}
	}
	l = 0
	for _, cnt := range cnts {
		if cnt > l {
			l = cnt
		}
	}
	ar := make([][]string, l)
	for s, cnt := range cnts {
		ar[cnt-1] = append(ar[cnt-1], s)
	}

	o, i := make([]string, len(cnts)), 0
	for _, ss := range ar {
		sort.Strings(ss)
		for _, s := range ss {
			o[i] = s
			i++
		}
	}
	return o
}
