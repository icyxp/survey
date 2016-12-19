package main

type Input struct {
	Message string
}

func (p *Input) Prompt() string {
	return p.Message
}

func (i *Input) HandleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	return line, pos, true
}
