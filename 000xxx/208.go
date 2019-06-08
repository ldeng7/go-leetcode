type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

type Trie struct {
	root TrieNode
}

func Constructor() Trie {
	return Trie{TrieNode{}}
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) find(key string) *TrieNode {
	n := &this.root
	for i := 0; i < len(key); i++ {
		j := key[i] - 'a'
		if nil == n.children[j] {
			return nil
		}
		n = n.children[j]
	}
	return n
}

func (this *Trie) Search(word string) bool {
	if n := this.find(word); nil != n {
		return n.isWord
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return nil != this.find(prefix)
}
