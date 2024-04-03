package structures

import "io"

type stack struct {
	eles []any
}

func (s *stack) evaluate(reader io.Reader) string {
	return evaluate(s, reader)
}

func (s *stack) elements() []any {
	return s.eles
}

func (s *stack) push(value any) {
	s.eles = append(s.eles, value)
}

func (s *stack) peek() any {
	if len(s.eles) == 0 {
		return nil
	}
	return s.eles[len(s.eles)-1]
}

func (s *stack) pop() any {
	if len(s.eles) == 0 {
		return nil
	}
	ele := s.eles[len(s.eles)-1]
	s.eles = s.eles[:len(s.eles)-1]
	return ele
}
