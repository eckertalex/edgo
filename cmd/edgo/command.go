package main

import (
	"strings"
)

type Command struct {
	name string
	args []string
}

func parseCommand(input string) Command {
	cmd := strings.TrimSpace(input)

	if strings.HasPrefix(cmd, "/") {
		if cmd == "/" || cmd == "//" {
			return Command{name: "search", args: []string{}}
		}
		searchTerm := strings.Trim(cmd, "/")
		return Command{name: "search", args: []string{searchTerm}}
	} else if strings.HasPrefix(cmd, "?") {
		if cmd == "?" || cmd == "??" {
			return Command{name: "reverse-search", args: []string{}}
		}
		searchTerm := strings.Trim(cmd, "?")
		return Command{name: "reverse-search", args: []string{searchTerm}}
	}

	if strings.HasPrefix(cmd, "e !") {
		parts := strings.Split(cmd, "e !")
		if len(parts) == 2 {
			return Command{name: "e", args: []string{"!", parts[1]}}
		}
	}

	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return Command{}
	}

	name := parts[0]
	args := parts[1:]

	validCommands := map[string]bool{
		"w": true, "e": true, "P": true, "p": true, "q": true, "x": true,
		"Q": true, ".": true, "$": true, "+": true, "-": true,
	}

	if validCommands[name] {
		return Command{name: name, args: args}
	}

	return Command{}
}
