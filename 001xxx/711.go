import "sort"

func countPairs(deliciousness []int) int {
	o, l := 0, len(deliciousness)
	sort.Ints(deliciousness)
	for p, s := 21, 1; p >= 0; p, s = p-1, s<<1 {
		for i, j := 0, l-1; i < j; {
			di, dj := deliciousness[i], deliciousness[j]
			if t := di + dj; t == s {
				if di == dj {
					o = (o + (j-i+1)*(j-i)/2) % 1000000007
					break
				}
				ci, cj := 1, 1
				for ; deliciousness[i] == deliciousness[i+1]; i, ci = i+1, ci+1 {
				}
				for ; deliciousness[j] == deliciousness[j-1]; j, cj = j-1, cj+1 {
				}
				o = (o + ci*cj) % 1000000007
				i, j = i+1, j-1
			} else if t > s {
				j--
			} else {
				i++
			}
		}
	}
	return o
}
