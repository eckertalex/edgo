package main

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		input    string
		expected Command
	}{
		{"P", Command{name: "P", args: []string{}}},
		{"h", Command{name: "h", args: []string{}}},
		{"H", Command{name: "H", args: []string{}}},
		{"e", Command{name: "e", args: []string{}}},
		{"p", Command{name: "p", args: []string{}}},
		{"n", Command{name: "n", args: []string{}}},
		{".", Command{name: ".", args: []string{}}},
		{"$", Command{name: "$", args: []string{}}},
		{"q", Command{name: "q", args: []string{}}},
		{"Q", Command{name: "Q", args: []string{}}},

		{"w filename.txt", Command{name: "w", args: []string{"filename.txt"}}},
		{"e filename.txt", Command{name: "e", args: []string{"filename.txt"}}},
		{"w", Command{name: "w", args: []string{}}},
		{"e", Command{name: "e", args: []string{}}},

		{"e !ls -la", Command{name: "e", args: []string{"!", "ls -la"}}},
		{"e !echo 'Hello, World!'", Command{name: "e", args: []string{"!", "echo 'Hello, World!'"}}},

		{"/re/", Command{name: "search", args: []string{"re"}}},
		{"/re", Command{name: "search", args: []string{"re"}}},
		{"/", Command{name: "search", args: []string{}}},
		{"//", Command{name: "search", args: []string{}}},

		{"?re?", Command{name: "reverse-search", args: []string{"re"}}},
		{"?re", Command{name: "reverse-search", args: []string{"re"}}},
		{"?", Command{name: "reverse-search", args: []string{}}},
		{"??", Command{name: "reverse-search", args: []string{}}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseCommand(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseCommand(%q) = %+v; want %+v", tt.input, result, tt.expected)
			}
		})
	}
}
