// Implementing using stack
package stack

type Stack struct {
	items []int
}

func (s *Stack) New() *Stack {
	s.items = []int{}
	return s
}

func (s *Stack) Push(t int) {
	s.items = append(s.items, t)
}

func (s *Stack) Pop() int {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}