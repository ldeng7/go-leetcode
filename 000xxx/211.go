type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

type WordDictionary struct {
	root TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	n := &this.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if nil == n.children[j] {
			n.children[j] = &TrieNode{}
		}
		n = n.children[j]
	}
	n.isWord = true
}

func (this *WordDictionary) find(key string, i int, n *TrieNode) bool {
	if i == len(key) {
		return n.isWord
	}
	if key[i] == '.' {
		for _, c := range n.children {
			if nil != c && this.find(key, i+1, c) {
				return true
			}
		}
		return false
	} else {
		c := n.children[key[i]-'a']
		return nil != c && this.find(key, i+1, c)
	}
}

func (this *WordDictionary) Search(word string) bool {
	return this.find(word, 0, &this.root)
}
