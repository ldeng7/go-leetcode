type ArrayStack struct {
	arr []byte
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{[]byte{}}
}

func (s *ArrayStack) Top() (byte, bool) {
	if len(s.arr) != 0 {
		return s.arr[len(s.arr)-1], true
	}
	return 0, false
}

func (s *ArrayStack) Push(item byte) {
	s.arr = append(s.arr, item)
}

func (s *ArrayStack) Pop() (byte, bool) {
	if len(s.arr) != 0 {
		item := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return item, true
	}
	return 0, false
}

func removeDuplicates(S string) string {
	st := NewArrayStack()
	for _, c := range []byte(S) {
		if t, ok := st.Top(); ok && t == c {
			st.Pop()
		} else {
			st.Push(c)
		}
	}
	return string(st.arr)
}
