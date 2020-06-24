import "sort"

func popCountU16(u uint16) uint16 {
	u = (u & 0x5555) + ((u >> 1) & 0x5555)
	u = (u & 0x3333) + ((u >> 2) & 0x3333)
	u = (u & 0x0f0f) + ((u >> 4) & 0x0f0f)
	u = (u & 0x00ff) + ((u >> 8) & 0x00ff)
	return u
}

func sortByBits(arr []int) []int {
	m := make(map[int]uint16, len(arr))
	for _, a := range arr {
		m[a] = popCountU16(uint16(a))
	}
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		ca, cb := m[a], m[b]
		return ca < cb || (ca == cb && a < b)
	})
	return arr
}
