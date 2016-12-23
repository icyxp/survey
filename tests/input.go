package main

import (
	"github.com/icyxp/survey"
	"github.com/icyxp/survey/tests/util"
)

var table = []TestUtil.TestTableEntry{
	{
		"no default", &survey.Input{"Hello world", ""},
	},
	{
		"default", &survey.Input{"Hello world", "default"},
	},
}

func main() {
	TestUtil.RunTable(table)
}
