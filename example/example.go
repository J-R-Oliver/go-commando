package main

import (
	"fmt"

	"github.com/J-R-Oliver/go-commando"
)

func main() {
	argOptPrinter := func(arguments []string, options map[string]string) {
		fmt.Println("Arguments:")

		for i, a := range arguments {
			fmt.Printf("\tindex: %d, argument: %s\n", i, a)
		}

		fmt.Println("Options:")

		for k, v := range options {
			fmt.Printf("\tkey: %s, option: %s\n", k, v)
		}
	}

	program := commando.NewProgram()

	program.
		Name("file-splitter").
		Description("CLI to split file written in go.").
		Version("1.0.0").
		Option("i", "input", "input", "Input file", "./input.txt").
		Option("o", "output", "output", "Output file", "./output.txt").
		Action(argOptPrinter).
		Parse()
}
