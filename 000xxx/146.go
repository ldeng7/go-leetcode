type ListNode struct {
	Val1 int
	Val2 int
	Prev *ListNode
	Next *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func (l *List) PushFront(n *ListNode) {
	n.Prev = nil
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	n.Next = l.Head
	l.Head = n
}

func (l *List) PopFront() *ListNode {
	if l.Head == nil {
		return nil
	}
	h := l.Head
	if h.Next != nil {
		l.Head = h.Next
	} else {
		l.Head, l.Tail = nil, nil
	}
	return h
}

func (l *List) PopBack() *ListNode {
	if l.Tail == nil {
		return nil
	}
	r := l.Tail
	if r.Prev != nil {
		l.Tail = r.Prev
	} else {
		l.Head, l.Tail = nil, nil
	}
	return r
}

func (l *List) Erase(n *ListNode) {
	if l.Head == n {
		l.PopFront()
		return
	} else if l.Tail == n {
		l.PopBack()
		return
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

type LRUCache struct {
	c int
	m map[int]*ListNode
	l *List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*ListNode{}, &List{}}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}
	this.l.Erase(n)
	this.l.PushFront(n)
	return n.Val2
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.m[key]
	if !ok {
		n = &ListNode{Val1: key, Val2: value}
		if len(this.m) == this.c {
			mp := this.l.PopBack()
			delete(this.m, mp.Val1)
		}
	} else {
		n.Val2 = value
		this.l.Erase(n)
	}
	this.m[key] = n
	this.l.PushFront(n)
}
