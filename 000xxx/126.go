import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	o := [][]string{}
	d, s := map[string]bool{}, map[string]bool{}
	for _, w := range wordList {
		d[w] = true
	}
	q := [][]string{{beginWord}}
	lvl, ml := 1, math.MaxInt64
	for 0 != len(q) {
		ar := q[0]
		q = q[1:]
		if len(ar) > lvl {
			for w, _ := range s {
				delete(d, w)
			}
			s = map[string]bool{}
			lvl = len(ar)
			if lvl > ml {
				break
			}
		}
		e := ar[len(ar)-1]
		for i := 0; i < len(e); i++ {
			bs := []byte(e)
			for c := byte('a'); c <= 'z'; c++ {
				bs[i] = c
				sn := string(bs)
				if !d[sn] {
					continue
				}
				s[sn] = true
				ar1 := make([]string, len(ar))
				copy(ar1, ar)
				ar1 = append(ar1, sn)
				if sn == endWord {
					o, ml = append(o, ar1), lvl
				} else {
					q = append(q, ar1)
				}
			}
		}
	}
	return o
}
