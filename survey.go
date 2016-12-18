package main

import "fmt"

var survey = []Prompt{
	&Password{"Hello"},
}

type Prompt interface {
	Prompt(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool)
}

func main() {
	// grab the readline instance
	rl, err := GetReadline()
	if err != nil {
		panic(err)
	}

	// keep prompting the user
	for _, prompt := range survey {
		// make sure it uses the correct listener
		rl.Config.SetListener(prompt.Prompt)

		// get the next line from the user
		line, err := rl.Readline()

		// if there was an error
		if err != nil {
			break
		}
		fmt.Println("responded with", line)
	}
}
