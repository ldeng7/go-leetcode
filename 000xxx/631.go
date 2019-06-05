import (
	"fmt"
	"strconv"
	"strings"
)

type Excel struct {
	mat [][]int
	m   map[string][]string
}

func Constructor(H int, W byte) Excel {
	mat := make([][]int, H)
	for i := 0; i < H; i++ {
		mat[i] = make([]int, W-'A'+1)
	}
	return Excel{mat, map[string][]string{}}
}

func getKey(r int, c byte) string {
	return fmt.Sprintf("%c%d", c, r)
}

func (this *Excel) Set(r int, c byte, v int) {
	delete(this.m, getKey(r, c))
	this.mat[r-1][c-'A'] = v
}

func (this *Excel) Get(r int, c byte) int {
	k := getKey(r, c)
	if strs, ok := this.m[k]; ok {
		return this.Sum(r, c, strs)
	}
	return this.mat[r-1][c-'A']
}

func (this *Excel) Sum(r int, c byte, strs []string) int {
	out := 0
	for _, str := range strs {
		idx := strings.LastIndexByte(str, ':')
		if idx >= 0 {
			ib, _ := strconv.Atoi(str[1:idx])
			ie, _ := strconv.Atoi(str[idx+2:])
			jb, je := str[0]-'A', str[idx+1]-'A'
			for i := ib; i <= ie; i++ {
				for j := jb; j <= je; j++ {
					out += this.Get(i, j+'A')
				}
			}
		} else {
			i, _ := strconv.Atoi(str[1:])
			out += this.Get(i, str[0])
		}
	}
	this.m[getKey(r, c)] = strs
	return out
}
