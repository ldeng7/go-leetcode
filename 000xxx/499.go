import "math"

func dup(bs []byte) []byte {
	o := make([]byte, len(bs))
	copy(o, bs)
	return o
}

var dirs [4][2]int = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var ws [4]byte = [4]byte{'u', 'd', 'l', 'r'}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	maze[hole[0]][hole[1]] = 2
	m, n := len(maze), len(maze[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = math.MaxInt64
		}
	}
	t[ball[0]][ball[1]] = 0
	q := [][2]int{{ball[0], ball[1]}}
	mp := map[int][]byte{}

	for 0 != len(q) {
		h := q[0]
		q = q[1:]
		for i, dir := range dirs {
			y, x := h[0], h[1]
			t0, j := t[y][x], y*n+x
			for y >= 0 && y < m && x >= 0 && x < n && maze[y][x] == 0 {
				y, x, t0 = y+dir[0], x+dir[1], t0+1
			}
			if y != hole[0] || x != hole[1] {
				y, x, t0 = y-dir[0], x-dir[1], t0-1
			}
			if y == h[0] && x == h[1] {
				continue
			}
			p := dup(mp[j])
			p = append(p, ws[i])

			j1 := y*n + x
			if t[y][x] > t0 {
				t[y][x], mp[j1] = t0, p
				if maze[y][x] != 2 {
					q = append(q, [2]int{y, x})
				}
			} else if t[y][x] == t0 && string(mp[j1]) > string(mp[j]) {
				mp[j1] = p
				if maze[y][x] != 2 {
					q = append(q, [2]int{y, x})
				}
			}
		}
	}

	out := mp[hole[0]*n+hole[1]]
	if 0 == len(out) {
		return "impossible"
	}
	return string(out)
}
