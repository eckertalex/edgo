package main

import (
	"bufio"
	"fmt"
	"io"
)

type Editor struct {
	lastError        error
	reader           io.Reader
	writer           io.Writer
	buffer           *Buffer
	prompt           string
	filename         string
	mode             EditorMode
	isPromptShown    bool
	isFullErrorShown bool
	running          bool
}

func NewEditor(reader io.Reader, writer io.Writer, prompt string, isPromptShown bool) *Editor {
	return &Editor{
		reader:        reader,
		writer:        writer,
		prompt:        prompt,
		isPromptShown: isPromptShown,
		mode:          ModeCommand,
		buffer:        NewBuffer(),
	}
}

func (ed *Editor) printPrompt() {
	if ed.isPromptShown {
		fmt.Fprint(ed.writer, ed.prompt)
	}
}

func (ed *Editor) printError() {
	fmt.Fprintln(ed.writer, "?")
	if ed.isFullErrorShown {
		fmt.Fprintln(ed.writer, ed.lastError)
	}
}

func (ed *Editor) executeCommand(cmd Command) {
	var err error

	switch cmd.name {
	case "P":
		err = ed.cmdTogglePrompt()
	case "h":
		err = ed.cmdShowLastError()
	case "H":
		err = ed.cmdToggleShowFullError()
	case "e":
		err = ed.cmdRead(cmd.args[0])
	case "p":
		err = ed.cmdPrint(false)
	case "n":
		err = ed.cmdPrint(true)
	case ".":
		err = ed.cmdPrint(false)
	case "$":
		fmt.Fprintln(ed.writer, "TODO")
	case "w":
		fmt.Fprintln(ed.writer, "TODO")
	case "q":
		err = ed.cmdQuit()
	case "Q":
		err = ed.cmdQuit()
	default:
		err = ErrUnknownCommand
		ed.printError()
	}

	if err != nil {
		ed.lastError = err
	}
}

func (ed *Editor) Run() error {
	scanner := bufio.NewScanner(ed.reader)
	ed.running = true

	for ed.running {
		ed.printPrompt()

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("input error: %w", err)
			}
			break
		}

		cmd := parseCommand(scanner.Text())
		ed.executeCommand(cmd)
	}

	return nil
}
