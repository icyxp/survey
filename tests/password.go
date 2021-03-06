package main

import (
	"github.com/icyxp/survey"
	"github.com/icyxp/survey/tests/util"
)

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Password{"Please type your password:"},
	},
	{
		"please make sure paste works", &survey.Password{"Please paste your password:"},
	},
}

func main() {
	TestUtil.RunTable(table)
}
