import (
	"bytes"
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
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
	n.isWord = true
}

func replaceWords(dict []string, sentence string) string {
	trie := &Trie{}
	buf := &bytes.Buffer{}
	words := strings.Split(sentence, " ")
	for _, w := range dict {
		trie.Add(w)
	}
	for _, w := range words {
		buf.WriteByte(' ')
		bs, n, i := make([]byte, 0, 8), &trie.root, 0
		for ; i < len(w); i++ {
			c := w[i]
			n = n.children[c-'a']
			if nil == n {
				buf.WriteString(w)
				break
			}
			bs = append(bs, c)
			if n.isWord {
				buf.Write(bs)
				break
			}
		}
		if i == len(w) {
			buf.WriteString(w)
		}
	}
	buf.ReadByte()
	return buf.String()
}
