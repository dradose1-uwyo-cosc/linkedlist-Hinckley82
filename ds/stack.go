package ds

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(value string) { // add new Head node
	s.list.InsertAt(0, value)
}

func (s *Stack) Pop() (string, bool) { // remove the Head
	if s.list.IsEmpty() {
		return "", false
	}

	value := s.list.Head.data
	s.list.RemoveAt(0)

	return value, true
}
