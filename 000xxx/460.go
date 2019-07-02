type listElemType = int

type ListNode struct {
	Val  listElemType
	list *List
	prev *ListNode
	next *ListNode
}

type List struct {
	root ListNode
	cnt  int
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *List) Len() int {
	return l.cnt
}

func (l *List) Head() *ListNode {
	if l.cnt != 0 {
		return l.root.next
	}
	return nil
}

func (l *List) Tail() *ListNode {
	if l.cnt != 0 {
		return l.root.prev
	}
	return nil
}

func (l *List) insertAfter(n, mark *ListNode) {
	next := mark.next
	n.prev, n.next, n.list = mark, next, l
	mark.next, next.prev = n, n
	l.cnt++
}

func (l *List) PushBack(n *ListNode) {
	if nil == n.list {
		l.insertAfter(n, l.root.prev)
	}
}

func (l *List) unlink(n *ListNode) {
	n.prev.next, n.next.prev = n.next, n.prev
	n.next, n.prev, n.list = nil, nil, nil
	l.cnt--
}

func (l *List) Unlink(n *ListNode) {
	if n.list == l {
		l.unlink(n)
	}
}

func (l *List) PopFront() *ListNode {
	if l.cnt != 0 {
		n := l.root.next
		l.unlink(n)
		return n
	}
	return nil
}

type LFUCache struct {
	capa, mf int
	m        map[int][]int
	f        map[int]*List
	n        map[int]*ListNode
}

func Constructor(capacity int) LFUCache {
	f := map[int]*List{1: (&List{}).Init()}
	return LFUCache{capacity, 0, map[int][]int{}, f, map[int]*ListNode{}}
}

func (this *LFUCache) Get(key int) int {
	p, ok := this.m[key]
	if !ok {
		return -1
	}
	c := p[1]
	this.f[c].Unlink(this.n[key])
	c++
	p[1] = c
	fn := this.f[c]
	if nil == fn {
		fn = (&List{}).Init()
		this.f[c] = fn
	}
	fn.PushBack(&ListNode{Val: key})
	if this.f[this.mf].Len() == 0 {
		this.mf++
	}
	this.n[key] = fn.Tail()
	return p[0]
}

func (this *LFUCache) Put(key int, value int) {
	if this.capa <= 0 {
		return
	}
	if this.Get(key) != -1 {
		this.m[key][0] = value
		return
	}
	if len(this.m) == this.capa {
		k := this.f[this.mf].Head().Val
		delete(this.m, k)
		delete(this.n, k)
		this.f[this.mf].PopFront()
	}
	this.m[key] = []int{value, 1}
	this.f[1].PushBack(&ListNode{Val: key})
	this.n[key] = this.f[1].Tail()
	this.mf = 1
}
