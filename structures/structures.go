package structures

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type pushPopStruct interface {
	push(value any)
	pop() any
	peek() any
	evaluate(reader io.Reader) string
}

// evaluate assumes the first line is the number of lines
// in the program and the rest are push, pop, peek commands
// anything else is ignored.
func evaluate(s pushPopStruct, reader io.Reader) string {
	var output []string
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	programLength, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < programLength; i++ {
		scanner.Scan()
		line := scanner.Text()
		switch line {
		case "pop":
			_ = s.pop()
		case "peek":
			output = append(output, fmt.Sprintf("%v", s.peek()))
		default:
			if len(line) < 5 {
				continue
			}
			s.push(line[5:])
		}
	}

	output = append(output, fmt.Sprintf("%v", s))
	return strings.Join(output, " ")

}
