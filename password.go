package main

type Password struct {
	Message string
}

func (p *Password) Prompt(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	newLine = append(line, '*')

	return newLine, pos, true
}
