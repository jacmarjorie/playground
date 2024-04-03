package structures

import (
	"io"
)

type queue struct {
	eles []any
}

func (q *queue) evaluate(reader io.Reader) string {
	return evaluate(q, reader)
}

func (q *queue) elements() []any {
	return q.eles
}

func (q *queue) push(value any) {
	q.eles = append(q.eles, value)
}

func (q *queue) peek() any {
	if len(q.eles) == 0 {
		return nil
	}
	return q.eles[0]
}

func (q *queue) pop() any {
	if len(q.eles) == 0 {
		return nil
	}
	ele := q.eles[0]
	q.eles = q.eles[1:]
	return ele
}

// rotate the array counter clockwise so
// that zero ends up at the front
// adding the counter give o(n)
func rotateLeftTillZero(q *queue) {
	cnt := len(q.eles)
	for q.peek() != 0 && cnt > 0 {
		cnt -= 1
		q.push(q.pop())
	}
}
