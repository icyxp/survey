package main

import (
	"github.com/icyxp/survey"
	"github.com/icyxp/survey/tests/util"
)

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Choice{
			Message: "Choose a color:",
			Choices: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
	{
		"short", &survey.Choice{
			Message: "Choose a color:",
			Choices: []string{"red", "blue"},
			Default: "red",
		},
	},
	{
		"default", &survey.Choice{
			Message: "Choose a color (should default blue):",
			Choices: []string{"red", "blue", "green"},
			Default: "blue",
		},
	},
}

func main() {
	TestUtil.RunTable(table)
}
