package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	promptFlag := flag.String("p", "", "Specify a command prompt. This may be toggled on and off with the P command.")

	flag.Parse()

	isPromptShown := *promptFlag != ""
	prompt := *promptFlag
	if prompt == "" {
		prompt = "*"
	}

	ed := NewEditor(os.Stdin, os.Stdout, prompt, isPromptShown)

	if flag.NArg() == 1 {
		ed.cmdRead(flag.Arg(0))
	}

	if err := ed.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
