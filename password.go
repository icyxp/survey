package main

type Password struct {
	Message string
}

func (p *Password) Prompt() string {
	return p.Message
}

func (p *Password) HandleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	newLine = append(line, '*')
	return line, pos, true
}
