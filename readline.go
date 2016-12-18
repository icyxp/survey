package main

// this file defines the interface between readline and the rest of the package

import (
	"fmt"

	"github.com/chzyer/readline"
)

// LineReader handles the dirty details of terminal i/o allowing for ea
type LineReader struct {
	*readline.Instance
}

// listenWith changes the internal key input handler
func (r *LineReader) listenWith(l readline.Listener) {
	r.Config.Listener = l
}

// handleInput is called on every line
func (r *LineReader) handleInput(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	// handle special keys
	switch key {
	// if the user pressed the up arrow
	case readline.CharPrev:
		fmt.Println("up arrow")
	// if the user pressed the down arrow
	case readline.CharNext:
		fmt.Println("down arrow")
	}

	return line, pos, true
}

// GetReadline returns the readline instance with the correct configuration
func GetReadline() (*LineReader, error) {
	// create an instance
	rl, err := readline.NewEx(&readline.Config{
		InterruptPrompt: "^C",
		EOFPrompt:       "Goodbye",
		HistoryLimit:    -1,
	})
	if err != nil {
		return nil, err
	}
	// create the reader out of the readline intsance
	reader := &LineReader{rl}

	return reader, nil
}
