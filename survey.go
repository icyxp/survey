package main

type Prompt interface {
	Prompt() string
	HandleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool)
}

func main() {

	// a list of prompts to play with
	var survey = []Prompt{
		&Password{"Hello"},
	}

	// grab the readline instance
	rl, err := GetReadline()
	if err != nil {
		panic(err)
	}

	// foo := ""

	// keep prompting the user
	for _, prompt := range survey {
		// make sure it uses the correct listener
		rl.Config.SetListener(prompt.HandleInput)
		// and the correct prompt
		rl.SetPrompt(prompt.Prompt())
		// get the next line from the user
		_, err := rl.Readline()
		// clear the prompt
		rl.SetPrompt("")

		// if there was an error
		if err != nil {
			panic(err)
		}
		// foo = line
		// fmt.Print("responded with ", line)
	}

	// fmt.Println(foo)
}
