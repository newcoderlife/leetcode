package min_stack

import "container/list"

type MinStack struct {
	buffer  *list.List
	current int
}

func Constructor() MinStack {
	return MinStack{buffer: list.New()}
}

func (s *MinStack) Push(val int) {
	if s.buffer.Len() == 0 {
		s.current = val
	}

	s.buffer.PushBack(val - s.current)
	if val < s.current {
		s.current = val
	}
}

func (s *MinStack) Pop() {
	if s.buffer.Len() == 0 {
		return
	}
	v := s.buffer.Back().Value.(int)
	if v < 0 {
		s.current -= v // 计算曾经的最小值
	}
	s.buffer.Remove(s.buffer.Back())
}

func (s *MinStack) Top() int {
	v := s.buffer.Back().Value.(int)
	if v > 0 {
		return v + s.current
	}
	return +s.current
}

func (s *MinStack) GetMin() int {
	return s.current
}
