type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Len  int
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Len {
		return -1
	}
	n := this.Head
	for i := 1; i <= index; i++ {
		n = n.Next
	}
	return n.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val, this.Head}
	this.Len++
	this.Head = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val, nil}
	if 0 == this.Len {
		this.Len = 1
		this.Head = node
		return
	}
	n := this.Head
	for nil != n.Next {
		n = n.Next
	}
	n.Next = node
	this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Len {
		this.AddAtTail(val)
	} else if index > 0 && index < this.Len {
		node := &Node{val, nil}
		n := this.Head
		for i := 1; i < index; i++ {
			n = n.Next
		}
		node.Next = n.Next
		n.Next = node
		this.Len++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.Head = this.Head.Next
		this.Len--
	} else if index > 0 && index < this.Len {
		n := this.Head
		for i := 1; i < index; i++ {
			n = n.Next
		}
		n.Next = n.Next.Next
		this.Len--
	}
}
