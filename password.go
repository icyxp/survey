package main

import (
	"github.com/chzyer/readline"
)

type Password struct {
	Message string
}

func (p *Password) Prompt(rl *readline.Instance) (string, error) {
	return "", nil
}
