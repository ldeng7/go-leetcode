type ListNode struct {
	Val  int
	prev *ListNode
	next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) PushFront(n *ListNode) {
	n.prev = nil
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	n.next = l.head
	l.head = n
}

func (l *List) InsertBefore(n, before *ListNode) {
	if before != l.head {
		n.prev, n.next = before.prev, before
		n.prev.next, before.prev = n, n
	} else {
		n.prev, n.next = nil, before
		before.prev, l.head = n, n
	}
}

func (l *List) InsertAfter(n, after *ListNode) {
	if after != l.tail {
		n.prev, n.next = after, after.next
		n.next.prev, after.next = n, n
	} else {
		n.prev, n.next = after, nil
		after.next, l.tail = n, n
	}
}

func (l *List) PopFront() *ListNode {
	if l.head == nil {
		return nil
	}
	h := l.head
	if h.next != nil {
		l.head = h.next
	} else {
		l.head, l.tail = nil, nil
	}
	return h
}

func (l *List) PopBack() *ListNode {
	if l.tail == nil {
		return nil
	}
	r := l.tail
	if r.prev != nil {
		l.tail = r.prev
	} else {
		l.head, l.tail = nil, nil
	}
	return r
}

func (l *List) Unlink(n *ListNode) {
	if l.head == n {
		l.PopFront()
		return
	} else if l.tail == n {
		l.PopBack()
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

type AllOne struct {
	mkv map[string]int
	mvk map[int]map[string]bool
	mvn map[int]*ListNode
	lv  *List
}

func Constructor() AllOne {
	return AllOne{map[string]int{}, map[int]map[string]bool{}, map[int]*ListNode{}, &List{}}
}

func (this *AllOne) addKV(key string, val int, onInc bool) {
	if m, ok := this.mvk[val]; ok {
		m[key] = true
	} else {
		this.mvk[val] = map[string]bool{key: true}
	}

	if len(this.mvk[val]) == 1 {
		n := &ListNode{Val: val}
		this.mvn[val] = n
		if onInc {
			if val != 1 {
				this.lv.InsertAfter(n, this.mvn[val-1])
			} else {
				this.lv.PushFront(n)
			}
		} else {
			this.lv.InsertBefore(n, this.mvn[val+1])
		}
	}
}

func (this *AllOne) removeKV(key string, val int) {
	m := this.mvk[val]
	if len(m) != 1 {
		delete(m, key)
	} else {
		delete(this.mvk, val)
		this.lv.Unlink(this.mvn[val])
		delete(this.mvn, val)
	}
}

func (this *AllOne) Inc(key string) {
	this.mkv[key]++
	v := this.mkv[key]
	this.addKV(key, v, true)
	if v != 1 {
		this.removeKV(key, v-1)
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.mkv[key]; !ok {
		return
	}
	this.mkv[key]--
	v := this.mkv[key]
	if v != 0 {
		this.addKV(key, v, false)
	}
	this.removeKV(key, v+1)
}

func (this *AllOne) GetOneKey(n *ListNode) string {
	if nil == n {
		return ""
	}
	keys := this.mvk[n.Val]
	for k, _ := range keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMaxKey() string {
	return this.GetOneKey(this.lv.tail)
}

func (this *AllOne) GetMinKey() string {
	return this.GetOneKey(this.lv.head)
}
