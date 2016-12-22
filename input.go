package main

import (
	"github.com/chzyer/readline"
)

type Input struct {
	Message string
}

func (p *Input) Prompt(rl *readline.Instance) (line string, err error) {
	// make sure the prompt matches the expectation
	rl.SetPrompt("> ")
	// get the next line
	line, err = rl.Readline()
	// we're done
	return
}

func (i *Input) HandleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	return line, pos, true
}
