package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (ed *Editor) cmdQuit() error {
	ed.running = false
	return nil
}

func (ed *Editor) cmdTogglePrompt() error {
	ed.isPromptShown = !ed.isPromptShown
	return nil
}

func (ed *Editor) cmdShowLastError() error {
	if ed.lastError != nil {
		fmt.Fprintln(ed.writer, ed.lastError)
	}
	return nil
}

func (ed *Editor) cmdToggleShowFullError() error {
	var err error

	ed.isFullErrorShown = !ed.isFullErrorShown
	if ed.isFullErrorShown {
		err = ed.cmdShowLastError()
	}

	return err
}

func (ed *Editor) cmdRead(filename string) error {
	ed.filename = filepath.ToSlash(filename)

	_, err := os.Stat(ed.filename)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", filename)
		return nil
	}

	if err != nil {
		return fmt.Errorf("error stating file: %w", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytesRead, err := ed.buffer.ReadFrom(file)
	if err != nil {
		return err
	}
	fmt.Fprintf(ed.writer, "%d\n", bytesRead)

	return nil
}

func (ed *Editor) cmdPrint(showLineNumber bool) error {
	var err error

	currentLine := ed.buffer.Current()
	if showLineNumber {
		_, err = fmt.Fprintf(ed.writer, "%d\t%s\n", ed.buffer.index, currentLine)
	} else {
		_, err = fmt.Fprintf(ed.writer, "%s\n", currentLine)
	}

	return err
}
