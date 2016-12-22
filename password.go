package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

type Password struct {
	Message   string
	HideInput bool
}

func (p *Password) handleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	for range line {
		fmt.Print("*")
	}
	// fmt.Println(line)
	return []rune("*"), pos - 1, false
}

func (p *Password) Prompt(rl *readline.Instance) (line string, err error) {
	// a configuration for the password prompt
	config := rl.GenPasswordConfig()
	// use the right prompt (make sure there is an empty space after the prompt)
	config.Prompt = p.Message + " "
	// if we are not supposed to hide the user's input
	if !p.HideInput {
		// add an empty line between the prompt and the input
		fmt.Print(" ")
		// then add the custom listener to the prompt
		config.Listener = readline.FuncListener(p.handleInput)
	}
	// ask for the user's Password
	pass, err := rl.ReadPasswordWithConfig(config)
	// we're done here
	return string(pass), err
}
