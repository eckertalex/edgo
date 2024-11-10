package main

import "strings"

type Command struct {
	name string
}

func parseCommand(input string) Command {
	return Command{name: strings.TrimLeft(input, " ")}
}
