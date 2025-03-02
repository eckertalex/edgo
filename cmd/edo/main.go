package main

import (
	"flag"
	"fmt"
	"os"
)

var Version = "v0.0.0"

func main() {
	var (
		promptFlag     = flag.String("p", "", "Specify a command prompt. This may be toggled on and off with the P command.")
		displayVersion = flag.Bool("v", false, "Display version and exit")
	)

	flag.Parse()

	if *displayVersion {
		fmt.Printf("%s version %s\n", os.Args[0], Version)
		os.Exit(0)
	}

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
