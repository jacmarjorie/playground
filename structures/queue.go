package structures

import (
	"io"
	"math/rand"
	"time"
)

// queue implements pushPopStruct interface
// TODO deque
type queue []any

func (q *queue) evaluate(reader io.Reader) string {
	return evaluate(q, reader)
}

func (q *queue) push(value any) {
	*q = append(*q, value)
}

func (q *queue) peek() any {
	if len(*q) == 0 {
		return nil
	}
	tmp := *q
	return tmp[0]
}

func (q *queue) pop() any {
	if len(*q) == 0 {
		return nil
	}
	tmp := *q
	*q = tmp[1:]
	return tmp[0]
}

// rotate the array counter clockwise so
// that zero ends up at the front
// adding the counter give o(n)
func rotateLeftTillZero(q *queue) {
	cnt := len(*q)
	for q.peek() != 0 && cnt > 0 {
		cnt -= 1
		q.push(q.pop())
	}
}

func GenerateRandomQueue(size int) *queue {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	zeroLoc := rand.Intn(size)
	q := new(queue)
	for i := 1; i < size; i++ {
		if i == zeroLoc {
			q.push(0)
			continue
		}
		q.push(i)
	}
	return q
}
