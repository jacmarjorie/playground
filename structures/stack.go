package structures

import "io"

// stack implements pushPopStruct interface
type stack []any

func (s *stack) evaluate(reader io.Reader) string {
	return evaluate(s, reader)
}

func (s *stack) push(value any) {
	*s = append(*s, value)
}

func (s *stack) peek() any {
	if len(*s) == 0 {
		return nil
	}
	tmp := *s
	return tmp[len(tmp)-1]
}

func (s *stack) pop() any {
	if len(*s) == 0 {
		return nil
	}
	tmp := *s
	*s = tmp[:len(tmp)-1]
	return tmp[len(tmp)-1]
}
