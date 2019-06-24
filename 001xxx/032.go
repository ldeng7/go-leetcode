type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

type StreamChecker struct {
	root  TrieNode
	nodes []*TrieNode
}

func Constructor(words []string) StreamChecker {
	m := map[string]bool{}
	this := &StreamChecker{}
	for i := 0; i < len(words); i++ {
		w := words[i]
		if !m[w] {
			this.add(w)
			m[w] = true
		}
	}
	return *this
}

func (this *StreamChecker) add(word string) {
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

func (this *StreamChecker) Query(letter byte) bool {
	o := false
	nodes := []*TrieNode{}
	j := letter - 'a'
	for _, node := range this.nodes {
		child := node.children[j]
		if nil != child {
			if child.isWord {
				o = true
			}
			nodes = append(nodes, child)
		}
	}
	if this.root.children[j] != nil {
		child := this.root.children[j]
		if child.isWord {
			o = true
		}
		nodes = append(nodes, child)
	}
	this.nodes = nodes
	return o
}
