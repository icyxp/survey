package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

type Prompt interface {
	Prompt(*readline.Instance) (string, error)
}

func main() {

	// a list of prompts to play with
	var survey = []Prompt{
		&Password{"Please enter your password:", false},
	}

	// grab the readline instance
	rl, err := GetReadline()
	if err != nil {
		panic(err)
	}

	// keep prompting the user
	for _, prompt := range survey {
		// get the next line from the user
		line, err := prompt.Prompt(rl)
		// if there was an error
		if err != nil {
			panic(err)
		}

		fmt.Println("You responded with", line)

	}
	// now that we are done asking questions, we need to close the readline terminal
	rl.Close()

}
