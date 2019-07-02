type TrieNode struct {
	children [26]*TrieNode
	str      string
}

type Trie struct {
	root TrieNode
}

func (t *Trie) Add(word string) {
	n := &t.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if nil == n.children[j] {
			n.children[j] = &TrieNode{}
		}
		n = n.children[j]
	}
	n.str = word
}

var dirs = [4][4]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	o := []string{}
	if 0 == len(board) || 0 == len(board[0]) || 0 == len(words) {
		return o
	}
	m, n := len(board), len(board[0])
	v := make([][]bool, m)
	for i := 0; i < m; i++ {
		v[i] = make([]bool, n)
	}
	trie := &Trie{}
	for _, w := range words {
		trie.Add(w)
	}

	var cal func(*TrieNode, int, int)
	cal = func(node *TrieNode, i, j int) {
		if 0 != len(node.str) {
			o = append(o, node.str)
			node.str = ""
		}
		v[i][j] = true
		for _, d := range dirs {
			y, x := i+d[0], j+d[1]
			if y >= 0 && y < m && x >= 0 && x < n && !v[y][x] {
				if n1 := node.children[board[y][x]-'a']; nil != n1 {
					cal(n1, y, x)
				}
			}
		}
		v[i][j] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if n := trie.root.children[board[i][j]-'a']; nil != n {
				cal(n, i, j)
			}
		}
	}
	return o
}
