package day17

var testinput = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

type Direction int

const (
	left Direction = iota
	right
)

type JetStream struct {
	pattern []Direction
	at      int
}

func (js *JetStream) Next() Direction {
	v := js.pattern[js.at%len(js.pattern)]
	js.at++
	return v
}

func Parse(input string) JetStream {
	js := JetStream{}
	for _, c := range input {
		switch c {
		case '>':
			js.pattern = append(js.pattern, right)
		case '<':
			js.pattern = append(js.pattern, left)
		default:
			panic("unexpected input")
		}
	}
	return js
}
