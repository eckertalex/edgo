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

	editor := NewEditor(os.Stdin, os.Stdout, prompt, isPromptShown)
	if err := editor.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
