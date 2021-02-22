import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const MAX_STONE = 18

var dirs = [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
var dist [MAX_STONE][105][105]int
var t [MAX_STONE - 2][(1 << (MAX_STONE - 2)) + 5]int

func minimalSteps(maze []string) int {
	m, n := len(maze), len(maze[0])
	stones, mines := make([]int, 0, 40), make([]int, 0, MAX_STONE)
	var yb, xb, ye, xe int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch maze[i][j] {
			case 'S':
				yb, xb = i, j
			case 'T':
				ye, xe = i, j
			case 'O':
				stones = append(stones, (i<<8)|j)
			case 'M':
				mines = append(mines, (i<<8)|j)
			}
		}
	}

	lm := len(mines)
	mines = append(mines, (yb<<8)|xb, (ye<<8)|xe)
	for k := 0; k < lm+2; k++ {
		distK := &(dist[k])
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				(*distK)[i][j] = math.MaxInt32
			}
		}
		mine := mines[k]
		(*distK)[mine>>8][mine&0xff] = 0
		q := make([]int, 1, 64)
		q[0] = mine
		for 0 != len(q) {
			t := q[0]
			q = q[1:]
			y, x := t>>8, t&0xff
			for _, dir := range dirs {
				y1, x1 := y+dir[0], x+dir[1]
				if y1 < 0 || y1 >= m || x1 < 0 || x1 >= n || maze[y1][x1] == '#' {
					continue
				}
				if t := (*distK)[y][x] + 1; (*distK)[y1][x1] > t {
					(*distK)[y1][x1] = t
					q = append(q, (y1<<8)|x1)
				}
			}
		}
	}
	if lm == 0 {
		if dist[0][ye][xe] == math.MaxInt32 {
			return -1
		}
		return dist[0][ye][xe]
	}

	edge := [MAX_STONE][MAX_STONE]int{}
	for i := 0; i < MAX_STONE; i++ {
		for j := 0; j < MAX_STONE; j++ {
			edge[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < lm+2; i++ {
		for j := 0; j < lm+2; j++ {
			for _, s := range stones {
				y, x := s>>8, s&0xff
				edge[i][j] = min(edge[i][j], dist[i][y][x]+dist[j][y][x])
			}
		}
	}

	te := 1 << lm
	for i := 0; i < lm; i++ {
		ti := &(t[i])
		for j := 0; j < te; j++ {
			(*ti)[j] = math.MaxInt32
		}
		(*ti)[1<<i] = edge[lm][i]
	}
	for i := 1; i < te; i++ {
		for j := 0; j < lm; j++ {
			if (i & (1 << j)) == 0 {
				continue
			}
			for k := 0; k < lm; k++ {
				if (i & (1 << k)) != 0 {
					continue
				}
				t[k][i|(1<<k)] = min(t[k][i|(1<<k)], t[j][i]+edge[j][k])
			}
		}
	}

	o := math.MaxInt32
	for i := 0; i < lm; i++ {
		o = min(o, t[i][(1<<lm)-1]+dist[i][ye][xe])
	}
	if o == math.MaxInt32 {
		return -1
	}
	return o
}
