package main

type Input struct {
	Message string
}

func (i *Input) Prompt(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	return line, pos, true
}
